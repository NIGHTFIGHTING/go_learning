// g++ bucket_limit.cpp -std=c++11
#include<iostream>
#include<sys/time.h>
#include<unordered_map>

using namespace std;

// 漏斗限流
class BucketLimit {
public:
    BucketLimit(int rate, int bucket):rate_(rate), bucket_(bucket), cur_bucket_(0), last_time_(now_micro_seconds() / kMicroSecondsPerSecond) {
    }
    void refresh() {
        int now = now_micro_seconds() / kMicroSecondsPerSecond; // 当前时间，单位为s
        cout << "now:" << now << " last_time_:" << last_time_ << endl;
        int consumer = (now - last_time_) * rate_; // 距离上次请求，这段时间已经消费多少请求
        cur_bucket_ = max(0, cur_bucket_ - consumer); // 当前剩余水桶中水量 
        ++cur_bucket_; // 水量增加1
    }
    bool allow() {
        refresh();
        cout << "cur_bucket_:" << cur_bucket_ << " bucket_:" << bucket_ << endl;
        if(cur_bucket_ < bucket_) {
            return true;
        }
        return false;
    }
private:
    int rate_; // 单位：r/s，每秒多少个request，qps
    int bucket_; // 水桶容量，超过水桶容量请求会被拒绝
    int64_t last_time_; // 以秒为单位,保存上一次请求的时间戳
    int cur_bucket_; // 当前水桶中存在的水量

    // 获取当前时间，单位为us
    int64_t now_micro_seconds() {
        struct timeval tv;
        gettimeofday(&tv, NULL);
        int64_t seconds = tv.tv_sec;
        return seconds * kMicroSecondsPerSecond + tv.tv_usec;
    }
    static const int kMicroSecondsPerSecond = 1000 * 1000;
};

int main() {
    std::unordered_map<int, bool> request_allow;
    // 限流50r/s，bucket为100r
    BucketLimit bucket_limit(50, 100);
    // for循环速度很快，这个for循环执行肯定是<1s
    for(int i = 0; i < 1000; ++i) {
        if(bucket_limit.allow()) {
            request_allow[i] = true;
        } else {
            request_allow[i] = false;
        }
    }
    for(auto ite : request_allow) {
        cout << "i = " << ite.first << " allow = " << ite.second << endl;
    }
    return 0;
}
