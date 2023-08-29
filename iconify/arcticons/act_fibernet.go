package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ActFibernet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.098 23.744l-3.727-11.469L11.5 23.744m1.29-3.871h5.018m10.248 0h0c0 2.15-1.72 3.87-3.87 3.87h0c-2.15 0-3.87-1.72-3.87-3.87v-3.87c0-2.15 1.72-3.87 3.87-3.87h0c2.15 0 3.727 1.72 3.727 3.87h0m.989-3.728H36.5m-3.727 11.469V12.275M8 26.325h3.919M8 30.245h2.548M8 26.325v7.839"/><circle cx="13.327" cy="26.619" r=".774" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.327 28.971v5.193m1.408-3.234c0-1.077.882-1.959 1.96-1.959h0c1.078 0 1.96.882 1.96 1.96v1.274c0 1.077-.882 1.96-1.96 1.96h0a1.966 1.966 0 0 1-1.96-1.96m0 1.959v-7.839m8.701 6.859c-.294.588-.98.98-1.666.98h0a1.966 1.966 0 0 1-1.96-1.96v-1.273c0-1.078.883-1.96 1.96-1.96h0c1.078 0 1.96.882 1.96 1.96v.686h-3.92m5.197-.687c0-1.077.882-1.959 1.96-1.959h0m-1.96 0v5.193m6.969 0v-3.233c0-1.078-.882-1.96-1.96-1.96h0c-1.078 0-1.96.882-1.96 1.96v3.233m0-3.233v-1.96m8.857 4.213c-.294.588-.98.98-1.666.98h0a1.966 1.966 0 0 1-1.96-1.96v-1.273c0-1.078.882-1.96 1.96-1.96h0c1.078 0 1.96.882 1.96 1.96v.686h-3.92m5.733-4.312v6.859m-1.078-5.193H40"/>`),
		g.Group(children),
	)
}