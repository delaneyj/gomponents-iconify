package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YoutubePlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m711 872l484-250l-484-253v503zM896 10q168 0 324.5 4.5T1450 24l73 4q1 0 17 1.5t23 3t23.5 4.5t28.5 8t28 13t31 19.5t29 26.5q6 6 15.5 18.5t29 58.5t26.5 101q8 64 12.5 136.5T1792 532v176q1 145-18 290q-7 55-25 99.5t-32 61.5l-14 17q-14 15-29 26.5t-31 19t-28 12.5t-28.5 8t-24 4.5t-23 3t-16.5 1.5q-251 19-627 19q-207-2-359.5-6.5T336 1256l-49-4l-36-4q-36-5-54.5-10t-51-21t-56.5-41q-6-6-15.5-18.5t-29-58.5T18 998q-8-64-12.5-136.5T0 748V572q-1-145 18-290q7-55 25-99.5T75 121l14-17q14-15 29-26.5T149 58t28-13t28.5-8t23.5-4.5t23-3t17-1.5q251-18 627-18z"/>`),
		g.Group(children),
	)
}