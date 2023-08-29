package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<defs><path id="cifBi0" d="M.5.5h300v180H.5z"/></defs><g fill="none" fill-rule="evenodd"><path fill="#CE1126" fill-rule="nonzero" d="M.5.5h300l-300 180h300z"/><path fill="#1EB53A" fill-rule="nonzero" d="M.5.5v180l300-180v180z"/><mask id="cifBi1" fill="#fff"><use href="#cifBi0"/></mask><path fill="#000" fill-rule="nonzero" stroke="#FFF" stroke-width="22.5" d="m.5.5l300 180m0-180l-300 180" mask="url(#cifBi1)"/><circle cx="150.5" cy="90.5" r="51" fill="#FFF" fill-rule="nonzero"/><path fill="#1EB53A" fill-rule="nonzero" d="m140.108 58.1l3.464 6l-3.464 6h6.928l3.464 6l3.464-6h6.928l-3.464-6l3.464-6h-6.928l-3.464-6l-3.464 6z"/><path fill="#CE1126" fill-rule="nonzero" d="m141.979 59.18l2.84 4.92l-2.84 4.92h5.681l2.84 4.92l2.841-4.92h5.681l-2.841-4.92l2.841-4.92h-5.681l-2.841-4.92l-2.84 4.92z"/><path fill="#1EB53A" fill-rule="nonzero" d="m117.233 97.475l3.464 6l-3.464 6h6.928l3.464 6l3.464-6h6.928l-3.464-6l3.464-6h-6.928l-3.464-6l-3.464 6z"/><path fill="#CE1126" fill-rule="nonzero" d="m119.104 98.555l2.84 4.92l-2.84 4.92h5.681l2.84 4.92l2.841-4.92h5.681l-2.841-4.92l2.841-4.92h-5.681l-2.841-4.92l-2.84 4.92z"/><path fill="#1EB53A" fill-rule="nonzero" d="m163.184 97.731l3.464 6l-3.464 6h6.929l3.464 6l3.464-6h6.928l-3.464-6.001l3.464-5.999h-6.928l-3.464-6l-3.464 6z"/><path fill="#CE1126" fill-rule="nonzero" d="m165.055 98.811l2.841 4.92l-2.841 4.92h5.681l2.841 4.919l2.84-4.919h5.681l-2.84-4.92l2.84-4.92h-5.681l-2.84-4.92l-2.841 4.919z"/></g>`),
		g.Group(children),
	)
}