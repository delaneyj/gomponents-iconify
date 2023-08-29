package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PpeFaceShieldAltNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g fill="currentColor" clip-path="url(#healthiconsPpeFaceShieldAltNegative0)"><path d="M36.848 28.008a1.663 1.663 0 0 0 1.39-.744a1.629 1.629 0 0 0 .124-1.56c-1.62-3.595-4.347-7.63-4.347-7.63l-.193-3.167c.001-.232-.06-.571-.147-.907h1.346a4.89 4.89 0 0 1 3.51 1.501A5.26 5.26 0 0 1 40 19.163V36h-5.241l.154-7.992h1.935Z"/><path fill-rule="evenodd" d="M48 0H0v48h48V0ZM34.759 36l-.04 2H42V19.163a7.26 7.26 0 0 0-2.033-5.054A6.89 6.89 0 0 0 35.02 12H32.8c-1.833-3.15-4.652-4.837-8.856-5.673c-6.313-1.253-12.132 1.182-15.348 5.774H22V12h10.8c.203.349.394.715.573 1.1c.099.217.213.556.302.9H24v12.215c0 2.605 1.004 5.097 2.782 6.93C28.56 34.977 30.963 36 33.463 36h1.296ZM22 14.1H7.428a14.84 14.84 0 0 0-1.39 5.226H22V14.1Zm0 7.226H6a14.65 14.65 0 0 0 1.258 5.295a14.762 14.762 0 0 0 3.7 5.036V42h16.578v-4H30v-.55a11.412 11.412 0 0 1-4.653-2.913C23.201 32.325 22 29.331 22 26.215v-4.89Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPpeFaceShieldAltNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}