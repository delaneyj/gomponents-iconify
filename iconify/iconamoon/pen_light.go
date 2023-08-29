package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 3v-.75a.75.75 0 0 0-.75.75H3Zm15.293 11.293l-.683-.31a.75.75 0 0 0 .153.84l.53-.53ZM21 17l.53.53a.75.75 0 0 0 0-1.06L21 17Zm-4 4l-.53.53a.75.75 0 0 0 1.06 0L17 21Zm-2.707-2.707l.53-.53a.75.75 0 0 0-.84-.153l.31.683ZM11 2.25H3v1.5h8v-1.5ZM19.75 11A8.75 8.75 0 0 0 11 2.25v1.5A7.25 7.25 0 0 1 18.25 11h1.5Zm-.774 3.602c.498-1.1.774-2.32.774-3.602h-1.5a7.22 7.22 0 0 1-.64 2.984l1.366.618Zm-1.213.221l2.707 2.707l1.06-1.06l-2.707-2.707l-1.06 1.06Zm2.707 1.647l-4 4l1.06 1.06l4-4l-1.06-1.06Zm-2.94 4l-2.707-2.707l-1.06 1.06l2.707 2.707l1.06-1.06ZM11 19.75a8.721 8.721 0 0 0 3.602-.774l-.618-1.366a7.22 7.22 0 0 1-2.984.64v1.5ZM2.25 11A8.75 8.75 0 0 0 11 19.75v-1.5A7.25 7.25 0 0 1 3.75 11h-1.5Zm0-8v8h1.5V3h-1.5Z"/><circle cx="11" cy="11" r="1" stroke="currentColor" stroke-width="1.5" transform="rotate(-180 11 11)"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 3l7 7"/></g>`),
		g.Group(children),
	)
}