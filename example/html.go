package example

import(
  "github.com/octu0/yuv"
)

func Yuv420pToRgb(r, u, v []uint8, width, height, yStride, uvStride int) {
  yuv420p  := yuv.NewYUV420p(width, height, yStride, uvStride)
  rgbPlane := yuv420p.ConvertRGBA(d[0], d[1], d[2])

  fmt.Println(`
  <style>
  .container { display: flex; flex-direction: row; }
  .dot { width: 100%; width: 1px; }
  .dot:before { content: ""; display: block; padding-top: 1px; }
  </style>
  `)
  i := 0
  for y := 0; y < height; y += 1 {
    fmt.Println("<div class='container'>");
    for x := 0; x < width; x += 1 {
      rgba := rgbPlane[i]
      r, g, b := rgba.R, rgba.G, rgba.B
      fmt.Printf("<div class='dot' style='background-color:#%02x%02x%02x'></div>", r, g, b)
      i += 1
    }
    fmt.Println("</div>");
  }
}
