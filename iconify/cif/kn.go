package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<defs><path id="cifKn0" d="M.5.5h300v200H.5z"/></defs><defs><use href="#cifKn0" id="cifKn1"/></defs><g fill="none" fill-rule="evenodd"><path fill="#009E49" fill-rule="nonzero" d="M300.5.5H.5v200"/><path fill="#CE1126" fill-rule="nonzero" d="M.5 200.5h300V.5"/><mask id="cifKn2" fill="#fff"><use href="#cifKn0"/></mask><path fill="#000" fill-rule="nonzero" stroke="#FCD116" stroke-width="80" d="m.5 200.5l300-200" mask="url(#cifKn2)"/><mask id="cifKn3" fill="#fff"><use href="#cifKn1"/></mask><path fill="#000" fill-rule="nonzero" stroke="#000" stroke-width="60" d="m.5 200.5l300-200" mask="url(#cifKn3)"/><path fill="#FFF" fill-rule="nonzero" d="m180.544 70.072l20.326.82l5.501 19.585l7.062-19.078l20.326.82l-15.962-12.611l7.061-19.079l-16.927 11.284l-15.963-12.61l5.502 19.585zm-114 76l20.326.82l5.501 19.585l7.062-19.078l20.326.82l-15.962-12.611l7.061-19.079l-16.927 11.284l-15.963-12.61l5.502 19.585z"/></g>`),
		g.Group(children),
	)
}