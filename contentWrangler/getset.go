package contentWrangler

func Get_Example(ctx context.Context) string{
  return ctx.Get(example).(string)
}

func Set_Example(ctx context.Context, in string) {
  return ctx.Set(example,in)
}
