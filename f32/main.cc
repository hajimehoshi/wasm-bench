#include <chrono>
#include <cstdio>
#include <memory>
#include <vector>

std::vector<float> array() {
  return std::vector<float>(16);
}

float heavyTask() {
  std::vector<float> arr = std::move(array());
  float result = 0.0f;
  for (int i = 0; i < 10000000; i++) {
    float u0 = 0.0f;
    float v0 = 0.0f;
    float u1 = 1.0f;
    float v1 = 1.0f;

    arr[0] = u0;
    arr[1] = v0;
    arr[2] = u1;
    arr[3] = u1;

    arr[4] = u1;
    arr[5] = v0;
    arr[6] = u0;
    arr[7] = u1;

    arr[8] = u0;
    arr[9] = v1;
    arr[10] = u1;
    arr[11] = u0;

    arr[12] = u1;
    arr[13] = v1;
    arr[14] = u0;
    arr[15] = u0;

    result += arr[i % 16];
  }
  return result;
}

int main() {
  auto t1 = std::chrono::high_resolution_clock::now();
  float x = heavyTask();
  auto t2 = std::chrono::high_resolution_clock::now();

  printf("result: %f\n", x);
  printf("time [ns]: %lld\n", std::chrono::duration_cast<std::chrono::nanoseconds>(t2-t1).count());
  return 0;
}
