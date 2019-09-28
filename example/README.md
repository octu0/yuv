### Example Output

This is an example of converting [BigBuckBunny](https://peach.blender.org/) yuv420 data to HTML RGB representation
YUV data can be expressed in HTML by using the following code(html.go).

```
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
```

### HTML YUV

see [big_buck_bunny.yuv.html](http://htmlpreview.github.io/?https://github.com/octu0/yuv/blob/master/example/big_buck_bunny.yuv.html)

### Big Buck Bunny

[Big Buck Bunny](https://peach.blender.org/) is licensed under the [Creative Commons Attribution 3.0 license](http://creativecommons.org/licenses/by/3.0/).
(c) copyright 2008, Blender Foundation / www.bigbuckbunny.org
