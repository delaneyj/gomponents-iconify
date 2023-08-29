package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pandas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#130754" d="M48.697 15.176h12.25v25.437h-12.25zm0 52.251h12.25v25.436h-12.25z" color="#000"/><path fill="#ffca00" d="M48.697 48.037h12.25v12.001h-12.25z" color="#000"/><path fill="#130754" d="M29.017 36.087h12.25v84.552h-12.25zM67.97 88.414h12.25v25.436H67.97zm0-52.297h12.25v25.437H67.97z" color="#000"/><path fill="#e70488" d="M67.97 68.983h12.25v12.001H67.97z" color="#000"/><path fill="#130754" d="M87.238 8.55h12.25v84.552h-12.25z" color="#000"/>`),
		g.Group(children),
	)
}