package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellRingingTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M9.63 17.531c.612.611.211 1.658-.652 1.706a3.992 3.992 0 0 1-3.05-1.166a3.992 3.992 0 0 1-1.165-3.049c.046-.826 1.005-1.228 1.624-.726l.082.074l3.161 3.16zM20.071 3.929c.96.96 1.134 2.41.52 3.547l-.09.153l-.024.036a8.013 8.013 0 0 1-1.446 7.137l-.183.223l-.191.218l-2.073 2.072l-.08.112a3 3 0 0 0-.499 2.113l.035.201l.045.185c.264.952-.853 1.645-1.585 1.051l-.086-.078L3.101 9.586c-.727-.727-.017-1.945.973-1.671a3 3 0 0 0 2.5-.418l.116-.086l2.101-2.1a8 8 0 0 1 7.265-1.86l.278.071l.037-.023a3.003 3.003 0 0 1 3.432.192l.14.117l.128.12z"/></g>`),
		g.Group(children),
	)
}