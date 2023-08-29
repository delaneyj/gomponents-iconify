package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M405 0H85Q50 0 25 25T0 85v342q0 35 25 60t60 25h342q35 0 60-25t25-60V107q0-45-31-76T405 0zM43 85q0-17 12.5-29.5T85 43h320q21 0 37.5 11.5T465 85H247l-70 58q-23 17-23 49t23 49l70 58h218q-6 19-22.5 30.5T405 341H85q-20 0-42 11V85zm42 299h299v21H49q11-21 36-21zm384 43q0 17-12.5 29.5T427 469H85q-26 0-36-21h378v-66q22-6 42-19v64zm0-171H265l-60-47q-9-5-9-17t9-17l60-47h204v128zm-170-64q0 9-6.5 15t-15.5 6q-8 0-14.5-6t-6.5-15t6.5-15t14.5-6q9 0 15.5 6t6.5 15z"/>`),
		g.Group(children),
	)
}