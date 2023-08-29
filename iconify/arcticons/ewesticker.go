package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ewesticker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.895 24.753a2.654 2.654 0 0 0 2.23 1.213l.052-.004a2.65 2.65 0 1 0 3.466 3.994a2.647 2.647 0 0 0 2.323 3.935a3.504 3.504 0 0 0 2.703-1.511v1.263a2.657 2.657 0 0 0 2.657 2.657a2.8 2.8 0 0 0 2.754-2.594a2.818 2.818 0 0 0 2.755 2.721a3.043 3.043 0 0 0 2.869-2.937a3.02 3.02 0 0 0 2.87 2.377a2.657 2.657 0 0 0 2.657-2.657a4.285 4.285 0 0 0-.804-2.246a3.158 3.158 0 0 0 2.762 1.854a2.656 2.656 0 0 0 1.451-4.881a2.657 2.657 0 0 0 .553-5.253a2.656 2.656 0 0 0-.35-5.29q-.107.003-.213.014a2.643 2.643 0 0 0-3.251-3.766q.008-.08.012-.16a2.65 2.65 0 0 0-4.17-2.177c-.026-1.449-1.702-2.61-3.15-2.61a2.883 2.883 0 0 0-2.809 2.763"/><circle cx="22.493" cy="13.71" r="4.391" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><ellipse cx="13.614" cy="13.601" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4.391" ry="3.404"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.133 22.381a2.317 2.317 0 1 1-4.633 0m26.084 0a2.317 2.317 0 1 1-4.634 0m-1.53-.681a6.948 6.948 0 1 1-13.896 0m-6.024.681c.001-2.157 1.678-6.742 4.077-7.442m22.007 7.442a10.36 10.36 0 0 0-4.374-8.269m-.26 8.269c.127-3.935-.981-5.83-.982-5.815m-11.479-6.389a6.503 6.503 0 0 1 2.434-2.452a3.34 3.34 0 0 1 1.503-.37A4.273 4.273 0 0 1 21.98 9.26M11.524 21.7s-.184-3.702.312-4.965m4.813 23.818a31.228 31.228 0 0 1 2.766-6.914m.387 6.975a13.848 13.848 0 0 1 1.544-4.396m11.964 4.41a20.047 20.047 0 0 0-.68-4.884m3.862 4.901a18.558 18.558 0 0 0-.897-5.613M25.42 21.7s.116-3.543-.278-4.402"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.93 28.634s-1.763-.19-1.19-2.573"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.08 24.518s1.927 2.956 5.742.736m4.172-1.155c.023-.567-1.341-2.043-2.012-2.263"/><circle cx="20.675" cy="13.671" r=".75" fill="currentColor"/><circle cx="15.439" cy="13.71" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.547 9.28s3.136-2.31 5.604.411m3.324-3.141s2.35-1.777 4.82 1.923M10.133 22.381s-.369-3.63 1.703-5.645"/>`),
		g.Group(children),
	)
}