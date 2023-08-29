package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brightnessaltauto(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m884 419l140 93l-140 93q0 2-1.5 6t-1.5 5l74 152l-169 11q-1 1-3.5 4t-3.5 4l-11 168l-151-74q-2 0-6 1.5t-6 1.5l-93 140l-93-140q-2 0-6-1.5t-6-1.5l-151 74l-11-169q-1-1-4-3.5t-4-3.5L69 768l74-152q0-1-1.5-5t-1.5-6L0 512l140-93q0-2 1.5-6t2.5-5L69 256l169-11q1-1 3.5-4t3.5-4l11-168l151 74q2 0 6-1.5t6-2.5L512 0l93 139q2 1 6 2.5t6 1.5l151-75l11 170q1 1 4 3.5t4 3.5l168 11l-74 151q0 2 1.5 6t1.5 6zM512 256q-106 0-181 75t-75 181q0 88 54 157t138 91v8h128v-8q84-22 138-91t54-157q0-106-75-181t-181-75zm96 384q-13 0-22.5-9.5T576 608v-32H448v32q0 13-9.5 22.5T416 640t-22.5-9.5T384 608V448q0-27 18.5-45.5T448 384h128q26 0 45 18.5t19 45.5v160q0 13-9.5 22.5T608 640zm-32-192H448v64h128v-64z"/>`),
		g.Group(children),
	)
}