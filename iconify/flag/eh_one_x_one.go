package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EhOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M0 0h512v256H0z"/><path fill="#007a3d" d="M0 256h512v256H0z"/><path fill="#fff" d="M0 149.3h512v213.3H0z"/><path fill="#c4111b" d="m0 0l256 256L0 512Z"/><g stroke-width="1.7" transform="translate(-135 -6.5) scale(1.02539)"><circle cx="512" cy="256" r="68.3" fill="#c4111b"/><circle cx="537.6" cy="256" r="68.3" fill="#fff"/><path fill="#c4111b" d="m493.7 297.3l29-20.8l29 21.2l-10.8-34.2l29-21l-35.8-.2l-11-34l-11.3 33.9l-35.7-.1l28.7 21.2l-11.1 34z"/></g>`),
		g.Group(children),
	)
}