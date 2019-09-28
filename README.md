# yuv

YUV(I420/yuv420p) utils.

this util feature:
-Convert I420 / YUV420 planar format data to rgb array
-Convert YCrCb to RGB

[https://github.com/octu0/yuv/tree/master/example](examples directory) contains examples of converting yuv data to a format that can be expressed in HTML.

# Usage

Prepare video data in YUV format and execute as below.

```
func main(){
  info    := decoder.Decode("foo.yuv")
  yuv420p := yuv.NewYUV420p(info.width, info.height, info.yStride, info.uvStride)
  for {
    frame := decoder.Next()
    decodedFrame(yuv420p, frame.Y, frame.U, frame.V)
  }
}
func decodedFrame(yuv420p *yuv.YUV420p, yPlane []uint8, uPlane uint8, vPlane uint8) {
  rgb := yuv420p.ConvertRGBA(d[0], d[1], d[2])
  i   := 0
  for y := 0; y < height; y += 1 {
    for x := 0; x < width; x += 1 {
      rgba := rgb[i]
      r, g, b := rgba.R, rgba.G, rgba.B
      fmt.Printf("r=%02x g=%02x b=%02x", r, g, b)
    }
    fmt.Println()
  }
}
```

see more [https://github.com/octu0/yuv/tree/master/example](example).


# License

BSD, see LICENSE file for details.