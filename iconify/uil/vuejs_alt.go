package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VuejsAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.867 3.374a1 1 0 0 0-.866-.502l-4.97-.004h-3.652a1.002 1.002 0 0 0-.817.425l-.559.796l-.564-.798a.998.998 0 0 0-.816-.423H6.968l-4.973.026A1 1 0 0 0 1.137 4.4l10.02 17.106a1 1 0 0 0 .863.494a1 1 0 0 0 .864-.496l9.98-17.128a1.002 1.002 0 0 0 .003-1.002ZM10.105 4.868l1.085 1.533a.999.999 0 0 0 .816.423h.001a1.003 1.003 0 0 0 .817-.425L13.9 4.87l1.363-.002l-3.244 5.454l-3.275-5.454Zm1.912 14.15L3.74 4.885l2.67-.015l4.754 7.918a1 1 0 0 0 .857.485h.006a1 1 0 0 0 .857-.489l4.708-7.916l2.67.003Z"/>`),
		g.Group(children),
	)
}