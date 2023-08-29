package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandGrabCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilHandGrabCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M10.183 6.72a1.875 1.875 0 0 1 3.645-.448a1.875 1.875 0 0 1 2.849 1.603v.21a1.5 1.5 0 0 1 .5-.085h.375c1.035 0 1.875.84 1.875 1.875v4.56a.493.493 0 0 1 .004.065c0 2.136-.806 3.774-2.043 4.874C16.16 20.465 14.538 21 12.931 21c-3.264 0-5.394-2.187-6.923-4.733a46.567 46.567 0 0 1-.978-1.559c-1.132-1.866.239-4.205 2.397-4.208V8.375a1.875 1.875 0 0 1 2.756-1.655Zm-.006 1.655a.875.875 0 1 0-1.75 0V13a.5.5 0 0 1-1 0v-1.5c-1.403.003-2.257 1.51-1.542 2.69c.357.59.697 1.135.962 1.533l.006.01l.007.01C8.33 18.193 10.199 20 12.93 20c1.394 0 2.771-.465 3.794-1.374c1.001-.89 1.69-2.23 1.707-4.062a.496.496 0 0 1-.005-.064V9.875A.875.875 0 0 0 17.553 9h-.375a.5.5 0 0 0-.5.5v1a.5.5 0 1 1-1 0V7.875a.875.875 0 1 0-1.75 0V10a.5.5 0 1 1-1 0V6.875a.875.875 0 1 0-1.75 0V10.5a.5.5 0 0 1-1 0V8.375Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilHandGrabCircleFilled0)"/></g>`),
		g.Group(children),
	)
}