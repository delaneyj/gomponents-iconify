package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LrOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagLr1x10"><path fill-opacity=".7" d="M0 0h512v512H0z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagLr1x10)"><path fill="#fff" d="M0 0h767.9v512H0z"/><path fill="#006" d="M0 0h232.7v232.8H0z"/><path fill="#c00" d="M0 464.9h767.9V512H0z"/><path fill="#c00" d="M0 465.4h767.9V512H0zm0-92.9h767.9v46.2H0zm0-93.2h766V326H0zM232.7 0h535.1v46.5H232.7zm0 186h535.1v46.8H232.7zm0-92.7h535.1v46.5H232.7z"/><path fill="#fff" d="m166.3 177.5l-50.7-31l-50.4 31.3l18.7-50.9l-50.3-31.4l62.3-.4l19.3-50.7L135 95h62.3l-50.1 31.7l19.1 50.8z"/></g>`),
		g.Group(children),
	)
}