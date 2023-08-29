package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSPill0"><g fill="none"><path fill="#fff" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M5.615 22.585a3 3 0 0 1 0-4.242L18.343 5.615a3 3 0 0 1 4.242 0l19.8 19.799a3 3 0 0 1 0 4.242L29.657 42.384a3 3 0 0 1-4.243 0L5.615 22.585Z"/><circle cx="14.808" cy="20.465" r="2" fill="#000" transform="rotate(-45 14.808 20.465)"/><circle cx="23.293" cy="28.949" r="2" fill="#000" transform="rotate(-45 23.293 28.95)"/><circle cx="19.05" cy="24.707" r="2" fill="#000" transform="rotate(-45 19.05 24.707)"/><circle cx="27.536" cy="33.193" r="2" fill="#000" transform="rotate(-45 27.536 33.193)"/><circle cx="20.464" cy="14.807" r="2" fill="#000" transform="rotate(-45 20.464 14.807)"/><circle cx="28.95" cy="23.293" r="2" fill="#000" transform="rotate(-45 28.95 23.293)"/><circle cx="24.707" cy="19.051" r="2" fill="#000" transform="rotate(-45 24.707 19.05)"/><circle cx="33.193" cy="27.535" r="2" fill="#000" transform="rotate(-45 33.193 27.535)"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSPill0)"/>`),
		g.Group(children),
	)
}