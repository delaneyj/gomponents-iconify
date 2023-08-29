package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneDisabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.8 22.7l-5.55-5.55q-2.225 1.8-4.838 2.825T4.05 21q-.6 0-.825-.3T3 19.95V15.9q0-.35.225-.613t.575-.337l3.45-.7q.275-.05.688.075t.612.325L10.9 17q.45-.275.975-.625t.925-.675L1.4 4.3l1.4-1.4l18.4 18.4l-1.4 1.4Zm-2.725-8.375L15.65 12.9q.35-.425.737-.975t.638-1.025L14.6 8.45q-.2-.2-.275-.563T14.3 7.3l.65-3.5q.075-.35.337-.575T15.9 3h4.05q.45 0 .75.3t.3.75q0 2.75-1.038 5.375t-2.887 4.9Z"/>`),
		g.Group(children),
	)
}