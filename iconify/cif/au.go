package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Au(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<defs><path id="cifAu0" d="M.5.5h150v75H.5z"/><path id="cifAu1" d="M.5.5V38h175v37.5h-25L.5.5zm150 0h-75V88H.5V75.5l150-75z"/></defs><g fill="none" fill-rule="evenodd"><path fill="#00008B" fill-rule="nonzero" d="M.5.5h300v150H.5z"/><mask id="cifAu2" fill="#fff"><use href="#cifAu0"/></mask><path fill="#000" fill-rule="nonzero" stroke="#FFF" stroke-width="17.578" d="m.5.5l150 75m0-75l-150 75" mask="url(#cifAu2)"/><mask id="cifAu3" fill="#fff"><use href="#cifAu1"/></mask><path fill="#000" fill-rule="nonzero" stroke="red" stroke-width="9.375" d="m.5.5l150 75m0-75l-150 75" mask="url(#cifAu3)"/><path fill="#000" fill-rule="nonzero" stroke="#FFF" stroke-width="23.438" d="M75.5.5V88M.5 38h175"/><path fill="#000" fill-rule="nonzero" stroke="red" stroke-width="14.063" d="M75.5.5v80.273M.5 38h155.273"/><path fill="#00008B" fill-rule="nonzero" d="M.5 75.5h150V.5h50v100H.5z"/><path fill="#FFF" fill-rule="nonzero" d="m75.5 90.5l4.339 13.491l13.252-5.019l-7.842 11.803l12.187 7.232l-14.118 1.228l1.944 14.037L75.5 123l-9.762 10.272l1.944-14.037l-14.118-1.228l12.187-7.232l-7.842-11.803l13.252 5.019zm150 24.285l2.066 6.424l6.311-2.39l-3.734 5.621l5.803 3.444l-6.722.585l.926 6.684l-4.649-4.891l-4.649 4.891l.926-6.684l-6.722-.585l5.803-3.444l-3.734-5.621l6.31 2.39zM188 55.411l2.066 6.424l6.311-2.39l-3.734 5.621l5.803 3.444l-6.722.585l.926 6.684l-4.65-4.892l-4.649 4.891l.926-6.684l-6.723-.585l5.803-3.444l-3.734-5.621l6.311 2.39zm37.5-40.625l2.066 6.424l6.311-2.39l-3.734 5.62l5.803 3.444l-6.722.585l.926 6.684l-4.649-4.891l-4.649 4.891l.926-6.684l-6.722-.585l5.803-3.444l-3.734-5.62l6.31 2.39zm33.333 30.625l2.066 6.424l6.311-2.39l-3.734 5.621l5.803 3.444l-6.723.585l.926 6.684l-4.649-4.891l-4.649 4.891l.926-6.684l-6.723-.585l5.803-3.444l-3.734-5.621l6.311 2.39zM240.5 75.5l1.633 4.003l4.311.316l-3.302 2.79l1.032 4.198l-3.674-2.279l-3.674 2.279l1.032-4.198l-3.302-2.79l4.311-.316z"/></g>`),
		g.Group(children),
	)
}