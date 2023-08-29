package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmVsiOnVpcForRegulatedIndustries(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9 21H3c-1.103 0-2-.897-2-2v-6c0-1.103.897-2 2-2h6c1.103 0 2 .897 2 2v6c0 1.103-.897 2-2 2zm-6-8v6h6v-6H3zm13 17c-.362 0-.72-.014-1.076-.04c-4.352-.332-8.36-2.732-10.723-6.42l1.685-1.08a12.057 12.057 0 0 0 9.19 5.505c.305.023.613.035.924.035v2zm12 0h-8c-1.103 0-2-.897-2-2v-8c0-1.103.897-2 2-2h8c1.103 0 2 .897 2 2v8c0 1.103-.897 2-2 2zm-8-10v8h8v-8h-8zm7.302-8c.454 1.282.698 2.621.698 4h2c0-1.37-.199-2.708-.584-4h-2.114zM27 10h-3c-1.103 0-2-.897-2-2V5c0-1.103.897-2 2-2h3c1.103 0 2 .897 2 2v3c0 1.103-.897 2-2 2zm-3-5v3h3V5h-3zm-4-2.417A13.952 13.952 0 0 0 16 2A13.951 13.951 0 0 0 4.202 8.46l1.684 1.08A11.961 11.961 0 0 1 20 4.698V2.583z"/>`),
		g.Group(children),
	)
}