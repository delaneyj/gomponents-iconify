package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IntercepterNg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="25.759" cy="24" r="7.706" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.18 24a15.818 15.818 0 0 0 27.158 0m0 0a15.818 15.818 0 0 0-27.159 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.32 24a9.035 9.035 0 0 0-.355-2.484l4.294-2.861l-2.094-3.807l-5.723 1.406a15.211 15.211 0 0 0-3.205-2.102l.869-5.142l-4.76-1.46l-3.677 4.73c-.303-.013-.603-.035-.91-.035s-.607.022-.91.035l-3.678-4.73l-4.76 1.46l.869 5.142a15.211 15.211 0 0 0-3.205 2.102l-5.722-1.406l-2.094 3.807l4.294 2.86a8.86 8.86 0 0 0 0 4.97l-4.294 2.86l2.094 3.807l5.722-1.406a15.211 15.211 0 0 0 3.205 2.102l-.868 5.142l4.759 1.46l3.678-4.73c.302.013.602.034.91.034s.607-.021.91-.034l3.678 4.73l4.759-1.46l-.868-5.142a15.211 15.211 0 0 0 3.205-2.102l5.722 1.406l2.094-3.807l-4.294-2.86A9.035 9.035 0 0 0 41.32 24Z"/>`),
		g.Group(children),
	)
}