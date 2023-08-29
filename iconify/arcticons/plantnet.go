package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plantnet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.323 10.879c-10.359 7.394-7.347 20.446-7.347 20.446l2.247-4.255m5.1-16.191c4.245 11.998-5.707 20.962-5.707 20.962l.6-4.774M12.22 13.205c5.354.35 9.816 5.801 9.816 5.801a17.16 17.16 0 0 0-1.052 9.976c-4.96-4.343-9.56-8.847-8.765-15.777ZM2.948 29.918a45.41 45.41 0 0 0 20.367 7.203c-3.638-8.421-10.103-11.44-20.367-7.203Zm21.801 7.157c11.025-.559 20.292-7.012 20.303-7.075c-9.38-4.045-16.193-1.799-20.303 7.075Z"/>`),
		g.Group(children),
	)
}