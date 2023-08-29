package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundwaveCircleBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<defs><path id="solarSoundwaveCircleBoldDuotone0" d="M12 6.25a.75.75 0 0 1 .75.75v10a.75.75 0 0 1-1.5 0V7a.75.75 0 0 1 .75-.75Zm-5 2a.75.75 0 0 1 .75.75v6a.75.75 0 0 1-1.5 0V9A.75.75 0 0 1 7 8.25Zm10 1a.75.75 0 0 1 .75.75v4a.75.75 0 0 1-1.5 0v-4a.75.75 0 0 1 .75-.75Z"/></defs><g fill="currentColor"><g opacity=".5"><path fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2Zm.75 5a.75.75 0 0 0-1.5 0v10a.75.75 0 0 0 1.5 0V7Zm-5 2a.75.75 0 0 0-1.5 0v6a.75.75 0 0 0 1.5 0V9Zm10 1a.75.75 0 0 0-1.5 0v4a.75.75 0 0 0 1.5 0v-4Z" clip-rule="evenodd"/><use href="#solarSoundwaveCircleBoldDuotone0"/></g><use href="#solarSoundwaveCircleBoldDuotone0"/></g>`),
		g.Group(children),
	)
}