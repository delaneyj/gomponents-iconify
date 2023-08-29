package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KnOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagKn1x10"><path fill-opacity=".7" d="M151.7-.3h745.1v745H151.7z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagKn1x10)" transform="translate(-104.2 .2) scale(.68714)"><path fill="#ffe900" d="M-5.3 0h1073.5v744H-5.3z"/><path fill="#35a100" d="m-5.8 0l1.2 536.4L830.7-.4L-5.8 0z"/><path fill="#c70000" d="m1069.5 744l-1.9-557.7L225 744.5l844.5-.4z"/><path d="m-5.3 576.9l.7 167.9l182.3-.3L1068 147.6l-1-146L886.9 0L-5.4 576.9z"/><path fill="#fff" d="m818 269l-64.2-2.2l-25.3 60.2l-14.3-61.5l-64.2-2.2l55.4-35.7L691 166l48.5 39.4l55.3-35.9l-25.4 60.2zM417.5 529.6l-64.3-2.3l-25.2 60.2l-14.3-61.5l-64.3-2.2l55.4-35.8l-14.4-61.4l48.5 39.4l55.3-35.9l-25.3 60.1z"/></g>`),
		g.Group(children),
	)
}