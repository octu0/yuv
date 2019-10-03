package yuv

import(
  "math"
  "image"
  "image/color"
)

//
// Format
//----------------------
// I420
//  YYYYYYYY
//  UU
//  VV
// NV12 (todo)
//  YYYYYYYY
//  UVUV
//
// Stride
//----------------------
// I420(Y-Stride:4, UV-stride:2)
//  YYYYYYYY----
//  UU--
//  VV--
//
// y plane
//====================================================
//   00 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15
//    Y  Y  Y  Y  Y  Y  Y  Y  -  -  -  -  -  -  -  -
//   16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31
//    Y  Y  Y  Y  Y  Y  Y  Y  -  -  -  -  -  -  -  -
//   32 33 34 35 ....
//    Y  Y  Y  Y  Y  Y  Y  Y  -  -  -  -  -  -  -  -
//
// u plane
//====================================================
//   00 01 02 03 04 05
//    U  U  U  U  -  - 
//   06 07 08 09 10 11
//    U  U  U  U  -  -
//
// v plane
//====================================================
//   00 01 02 03 04 05
//    V  V  V  V  -  - 
//   06 07 08 09 10 11
//    V  V  V  V  -  -
//


// I420/YUV420 planar format
type YUV420p struct {
  Width    int
  Height   int
  YStride  int
  UVStride int
}
func NewYUV420p(w, h, ys, uvs int) *YUV420p {
  n         := new(YUV420p)
  n.Width    = w
  n.Height   = h
  n.YStride  = ys
  n.UVStride = uvs
  return n
}

// convert yuv420p to rgba
func (n *YUV420p) ConvertRGBA(yPlane, uPlane, vPlane []byte) []RGBA {
  yPos, uPos, vPos := 0, 0, 0
  i := 0
  rgbPlane := make([]RGBA, n.Width * n.Height)
  for h := 0; h < n.Height; h += 1 {
    for w := 0; w < n.Width; w += 1 {
      y := 0xff & yPlane[w + yPos]
      u := 0xff & uPlane[(w / 2) + uPos]
      v := 0xff & vPlane[(w / 2) + vPos]
      r, g, b := rgb(y, u, v)
      rgbPlane[i] = RGBA{
        R: r,
        G: g,
        B: b,
        A: 0xff, // unused
      }
      i += 1
    }
    yPos = yPos + n.YStride
    if 0 < h && 0 == (h & 1) {
      uPos += n.UVStride
      vPos += n.UVStride
    }
  }
  return rgbPlane
}
func ConvertRGBAtoYCbCr(src *image.RGBA) *image.YCbCr {
  rectangle := src.Bounds()
  img := image.NewYCbCr(rectangle, image.YCbCrSubsampleRatio420)
  for y := rectangle.Min.Y; y < rectangle.Max.Y; y += 1 {
    for x := rectangle.Min.X; x < rectangle.Max.X; x += 1 {
      rgba       := src.RGBAAt(x, y)
      yy, uu, vv := color.RGBToYCbCr(rgba.R, rgba.G, rgba.B)

      cy := img.YOffset(x, y)
      ci := img.COffset(x, y)
      img.Y[cy]  = yy
      img.Cb[ci] = uu
      img.Cr[ci] = vv
    }
  }
  return img
}

//
// R = Y + 1.402   * (Cr - 128)
// G = Y - 0.34414 * (Cb - 128) - 0.71414 * (Cr - 128)
// B = Y + 1.772   * (Cb - 128)
//
func rgb(y, u, v uint8) (r uint8, g uint8, b uint8) {
  yf, uf, vf := float64(y), float64(u), float64(v)
  rf := yf + 1.402   * (vf - 128)
  gf := yf - 0.34414 * (uf - 128) - (0.71414 * (vf - 128))
  bf := yf + 1.772   * (uf - 128)
  r   = uint8(clamp(rf, 0, 255))
  g   = uint8(clamp(gf, 0, 255))
  b   = uint8(clamp(bf, 0, 255))
  return
}
func clamp(value float64, min, max float64) float64 {
  return math.Min(math.Max(value, min), max)
}
