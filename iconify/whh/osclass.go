package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Osclass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 256q-56 0-125-32q61 104 61 224q0 91-35.5 174T765 765t-143 95.5T448 896t-174-35.5T131 765T35.5 622T0 448t35.5-174T131 131t143-95.5T448 0q152 0 274 94q29-26 64-53.5T832 6q16-9 32.5-4.5t25 20.5t4 34T874 83q-30 19-107 50q0 1 .5 1.5t1.5.5q54 10 122 27.5T992 192q32 11 32 32q0 13-19.5 22.5T960 256zm-768 0q-64 86-64 192q0 32 6 64h64q-6-17-6-32q0-40 28-68t68-28t68 28t28 68q0 15-6 32h140q-6-17-6-32q0-40 28-68t68-28t68 28t28 68q0 15-6 32h63q7-32 7-64q0-106-65-192H192z"/>`),
		g.Group(children),
	)
}