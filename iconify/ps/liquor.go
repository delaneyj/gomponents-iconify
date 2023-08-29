package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Liquor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M359 226q49-23 49-66q0-31-22-53t-53-22q-16 0-34 7q2-2 2-7V21q0-21-21-21H152q-21 0-21 21v64q0 22 21 22q0 27-7.5 53T122 194q-63 31-94.5 83T7 386q8 53 47.5 87t97.5 39h141q53-7 88.5-37t45.5-78q7-41-10-88t-58-83zm6-66q0 11-13 21.5T323 192v9q-2-1-5.5-3.5T312 194q-13-7-28-47q26-19 49-19q13 0 22.5 9.5T365 160zM173 43h86v21h-86V43zm212 347q-7 34-32.5 55T289 469H154q-41-3-70.5-27.5T47 380q-7-39 14.5-80t77.5-67q31-15 43.5-52t12.5-74h42q0 21 2 32q0 1-1 3t-1 3h2q15 67 52 88q53 28 77 72.5t17 84.5zM216 256q-54 0-91 28t-37 68t37 68t91 28t91-28t37-68t-37-68t-91-28zm0 149q-35 0-60-15.5T131 352t25-37.5t60-15.5t60 15.5t25 37.5t-25 37.5t-60 15.5z"/>`),
		g.Group(children),
	)
}