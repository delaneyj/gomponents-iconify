package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.94 6.412A2 2 0 0 0 2 8.108V16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8.108a2 2 0 0 0-.94-1.696l-6-3.75a2 2 0 0 0-2.12 0l-6 3.75Zm2.615 2.423a1 1 0 1 0-1.11 1.664l5 3.333a1 1 0 0 0 1.11 0l5-3.333a1 1 0 0 0-1.11-1.664L10 11.798L5.555 8.835Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}