#include <iostream>
#include <vector>
#include <algorithm>
#include <limits>
#include <functional>

using namespace std;

int max_covers(const vector<int>& stamps, int N) {
    int stamp_type = stamps.size();
    int max_value = (stamp_type > 0 ? stamps.back() : 0) * N;
    if (max_value == 0) return 0;

    vector<vector<int>> min_use(stamp_type, vector<int>(max_value + 1, numeric_limits<int>::max()));

    for (int j = 0; j <= max_value; ++j) {
        if (j > N) break;
        min_use[0][j] = j;
    }

    for (int i = 0; i < stamp_type; ++i) {
        min_use[i][0] = 0;
    }

    for (int i = 1; i < stamp_type; ++i) {
        for (int j = 1; j <= max_value; ++j) {
            min_use[i][j] = min_use[i-1][j];
            for (int t = 1; t <= N; ++t) {
                if (j < t * stamps[i]) break;
                if (j - t * stamps[i] >= 0 && min_use[i - 1][j - t * stamps[i]] != numeric_limits<int>::max()) {
                    min_use[i][j] = min(min_use[i][j], t + min_use[i - 1][j - t * stamps[i]]);
                }
            }
        }
    }

    int max_reach = 0;
    for (int j = 1; j <= max_value; ++j) {
        if (min_use[stamp_type - 1][j] <= N) {
            max_reach = j;
        } else {
            break;
        }
    }
    return max_reach;
}

int main() {
    int N, K;
    cin >> N >> K;

    vector<int> path = {1};
    int ans = 1;
    vector<int> stamps;

    function<void(int)> backtrace = [&](int i) {
        if (i == K) {
            int r = max_covers(path, N);
            if (r > ans) {
                ans = r;
                stamps = path;
            }
            return;
        }

        for (int v = path.back() + 1; v <= ans + 2; ++v) {
            path.push_back(v);
            backtrace(i + 1);
            path.pop_back();
        }
    };

    backtrace(1);

    for (size_t i = 0; i < stamps.size(); ++i) {
        cout << stamps[i] << (i == stamps.size() - 1 ? "" : " ");
    }
    cout << endl;
    cout << "MAX=" << ans << endl;

    return 0;
}