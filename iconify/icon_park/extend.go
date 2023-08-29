package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Extend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" stroke-width="4" rx="3"/><path fill="#fff" d="M35 12H30.4142C29.5233 12 29.0771 13.0771 29.7071 13.7071L34.2929 18.2929C34.9229 18.9229 36 18.4767 36 17.5858V13C36 12.4477 35.5523 12 35 12Z"/><path fill="#fff" d="M12 13V17.5858C12 18.4767 13.0771 18.9229 13.7071 18.2929L18.2929 13.7071C18.9229 13.0771 18.4767 12 17.5858 12H13C12.4477 12 12 12.4477 12 13Z"/><path fill="#fff" d="M13 36H17.5858C18.4767 36 18.9229 34.9229 18.2929 34.2929L13.7071 29.7071C13.0771 29.0771 12 29.5233 12 30.4142V35C12 35.5523 12.4477 36 13 36Z"/><path fill="#fff" d="M36 35V30.4142C36 29.5233 34.9229 29.0771 34.2929 29.7071L29.7071 34.2929C29.0771 34.9229 29.5233 36 30.4142 36H35C35.5523 36 36 35.5523 36 35Z"/></g>`),
		g.Group(children),
	)
}