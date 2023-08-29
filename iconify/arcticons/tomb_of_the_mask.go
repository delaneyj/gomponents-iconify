package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TombOfTheMask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.254 5.5H18.171v19.278h3.812v3.381h14.459v-3.381h3.812V5.5zM25.4 22.98h7.625m-20.256-6.151h5.402m-5.402-6.258h5.402m17.803 6.438v-3.524M25.4 17.009v-3.524m-7.229 17.841h18.27V42.5h-18.27zm9.135 4.624v6.55M12.434 24.778v-3.02m0 20.742v-5.812m5.737-11.91H7.746v11.91h4.688v-5.362h5.737v-6.548z"/>`),
		g.Group(children),
	)
}