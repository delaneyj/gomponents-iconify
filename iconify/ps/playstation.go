package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Playstation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m250 371l-71-30V13q6 1 16.5 3t38 10T283 43t43 25t27 34q10 32 6 55.5T346 194t-16 14q-14 2-25-2.5t-15-9.5l-4-5V89l-35-15zm18-45v38q166-53 190-74q16-14-19-38q-42-30-78-25q-2 1-4 1q-46 3-88 21v35l106-26l38 17zm-106-3q-22 11-70 11q-39-1-65-12T2 291q0-21 48-42.5T161 217v40l-85 23q-14 18 5 19q20 0 40.5-2.5T152 291l10-3v35z"/>`),
		g.Group(children),
	)
}