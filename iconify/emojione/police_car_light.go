package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PoliceCarLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#3e4347" d="M9.3 52.5v4.2C9.3 59.1 19.5 61 32 61s22.7-1.9 22.7-4.2v-4.2l-45.4-.1"/><path fill="#62696d" d="M12.8 52.5v4.2c0 2.3 8.6 4.2 19.2 4.2s19.2-1.9 19.2-4.2v-4.2H12.8"/><path fill="#c94747" d="M54.7 52.5H9.3l5-25.1h35.4z"/><ellipse cx="32" cy="52.5" fill="#b23838" rx="22.7" ry="4.2"/><path fill="#ffc7ce" d="M51.2 52.2L47 27.4H17l-4.2 24.9v.3c0 2.3 8.6 4.2 19.2 4.2s19.2-1.9 19.2-4.2v-.4" opacity=".3"/><ellipse cx="32" cy="27.4" fill="#ed4c5c" rx="17.7" ry="4.2"/><path fill="#c94747" d="M17.2 27.8c0-.7.6-1.2 1.2-1.5c.6-.4 1.2-.6 1.8-.9c1.3-.4 2.6-.7 3.9-1c2.6-.5 5.3-.6 7.9-.7c2.6 0 5.3.2 7.9.7c1.3.2 2.6.5 3.9 1c.6.2 1.2.5 1.8.9c.5.3 1.2.8 1.2 1.5c-.1-.7-.8-1-1.3-1.3c-.6-.3-1.2-.5-1.8-.6c-1.3-.3-2.5-.5-3.8-.7c-2.6-.3-5.2-.4-7.8-.4c-2.6 0-5.2.1-7.8.4c-1.3.1-2.6.4-3.8.7c-.6.2-1.3.3-1.8.6c-.7.3-1.4.7-1.5 1.3" opacity=".5"/><ellipse cx="32" cy="40.2" fill="#ffc7ce" opacity=".3" rx="13.5" ry="13.4"/><path fill="#ffe6ea" d="m44.7 40.2l-7.8-2l4.3-7.2l-7.2 4.4l-2-7.9l-2 7.9l-7.2-4.4l4.3 7.2l-7.8 2l7.8 2l-4.3 7.2l7.2-4.3l2 7.8l2-7.8l7.2 4.3l-4.3-7.2z"/><path fill="#f15744" d="M56.9 8L49 19.9l12.8-7zM7.1 8L15 19.9l-12.8-7zM0 23.1l12 4l-12 1.7zm64 0l-12 4l12 1.7zM17.2 4l8.1 14.2L23.7 3zm29.6 0l-8.1 14.2L40.3 3z"/>`),
		g.Group(children),
	)
}