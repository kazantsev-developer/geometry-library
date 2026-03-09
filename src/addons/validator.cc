#include <napi.h>

bool ValidateTriangleInequality(double a, double b, double c) {
  return (a + b > c) && (a + c > b) && (b + c > a);
}

Napi::Boolean ValidateTriangle(const Napi::CallbackInfo &info) {
  Napi::Env env = info.Env();

  if (info.Length() < 3 || !info[0].IsNumber() || !info[1].IsNumber() ||
      !info[2].IsNumber()) {
    Napi::TypeError::New(env, "Arguments must be numbers")
        .ThrowAsJavaScriptException();
    return Napi::Boolean::New(env, false);
  }

  double a = info[0].As<Napi::Number>().DoubleValue();
  double b = info[1].As<Napi::Number>().DoubleValue();
  double c = info[2].As<Napi::Number>().DoubleValue();

  if (a <= 0 || b <= 0 || c <= 0) {
    return Napi::Boolean::New(env, false);
  }

  return Napi::Boolean::New(env, ValidateTriangleInequality(a, b, c));
}

Napi::Object Init(Napi::Env env, Napi::Object exports) {
  exports.Set("validateTriangle", Napi::Function::New(env, ValidateTriangle));
  return exports;
}

NODE_API_MODULE(validator, Init)