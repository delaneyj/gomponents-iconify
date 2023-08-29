package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Na(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<defs><path id="cifNa0" d="M.5.5h300v200H.5z"/></defs><defs><use href="#cifNa0" id="cifNa1"/></defs><g fill="none" fill-rule="evenodd"><path fill="#003580" fill-rule="nonzero" d="M300.5.5H.5v200z"/><path fill="#009543" fill-rule="nonzero" d="M.5 200.5h300V.5z"/><mask id="cifNa2" fill="#fff"><use href="#cifNa0"/></mask><path fill="#000" fill-rule="nonzero" stroke="#FFF" stroke-width="60" d="m.5 200.5l300-200" mask="url(#cifNa2)"/><mask id="cifNa3" fill="#fff"><use href="#cifNa1"/></mask><path fill="#000" fill-rule="nonzero" stroke="#D21034" stroke-width="40" d="m.5 200.5l300-200" mask="url(#cifNa3)"/><path fill="#FFCE00" fill-rule="nonzero" d="m60.5 22.166l5.176 14.015l11.49-9.549l-2.525 14.726l14.726-2.525l-9.549 11.49l14.015 5.176l-14.015 5.176l9.549 11.49l-14.726-2.524l2.525 14.725l-11.49-9.549L60.5 88.832l-5.176-14.015l-11.49 9.549l2.524-14.725l-14.725 2.524l9.549-11.49l-14.015-5.176l14.015-5.176l-9.549-11.49l14.725 2.525l-2.524-14.726l11.49 9.549z"/><circle cx="60.5" cy="55.5" r="18.333" fill="#FFCE00" fill-rule="nonzero" stroke="#003580" stroke-width="2"/></g>`),
		g.Group(children),
	)
}