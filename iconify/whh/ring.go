package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ring(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 576q0 91-35.5 174T765 893t-143 95.5t-174 35.5t-174-35.5T131 893T35.5 750T0 576q0-163 103.5-286.5T364 136L256 64L384 0h128l128 64l-108 72q157 30 260.5 153.5T896 576zM448 256q-87 0-160.5 43T171 415.5T128 576t43 160.5T287.5 853T448 896t160.5-43T725 736.5T768 576t-43-160.5T608.5 299T448 256z"/>`),
		g.Group(children),
	)
}