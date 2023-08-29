package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M13.223 59.348h45.745V32.776H13.223v26.572z"/><path fill="#9B9B9A" d="m18.74 38.284l36.48.177l.407-2.182H16.624v11.169l2.082.336l.034-9.5z"/><path fill="#d0cfce" d="M59.968 23.793H12.223v2.828l47.745-.015z"/><circle cx="22.268" cy="13.569" r="1" fill="#d0cfce"/><circle cx="22.49" cy="41.984" r="2" fill="#3f3f3f"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="m47.434 27.214l7.785 4.214M24.028 14.546l10.637 5.757"/><circle cx="22.267" cy="13.569" r="2"/><path d="M12.223 31.777h47.745v28.571H12.223zm4.323 20.46h38.572m-38.572 4h38.572"/><path d="M16.624 36.279h38.595v11.504H16.624z"/><circle cx="22.491" cy="41.985" r="2"/><path d="M12.223 26.946v-3.153h47.745v3.153M38.294 41.697v.667m4-.667v.667m4-.667v.667"/></g>`),
		g.Group(children),
	)
}