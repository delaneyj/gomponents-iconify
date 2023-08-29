package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParachuteSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M960 0q95 0 188 16t183 48q123 44 230 119t187 174t126 219t46 256q0 50-30 88t-69 65q-5 3-14 8t-20 11t-18 10t-11 7l-478 478v453q0 40-28 68t-68 28H736q-40 0-68-28t-28-68v-453l-478-478q-2-2-10-7t-19-10t-19-10t-15-9q-39-27-69-65T0 832q0-135 46-255t126-220t187-174T589 64q90-32 183-48T960 0zm-64 1151q-119-2-236-13t-235-35l306 305h165v-257zm293 257l306-305q-117 23-234 34t-237 14v257h165z"/>`),
		g.Group(children),
	)
}