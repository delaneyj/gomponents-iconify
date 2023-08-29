package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pdiddy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 171q-10-73-64.5-122T256 0T128.5 49.5T66 171h-2q-27 0-45.5 18T0 235q0 33 19 60t49 38q12 72 64.5 125.5T256 512t123.5-53.5T444 333q29-12 48.5-39t19.5-59q0-28-18.5-46T448 171zM151 85q107 40 210 0q36 34 42 86H109q6-49 42-86zM64 213v69q-21-18-21-47q0-22 21-22zm192 256q-55 0-97-45.5T109 320h74q30 0 41-30l6-36q5-21 26-21t26 21l6 36q11 30 41 30h74q-8 58-50 103.5T256 469zm192-187v-69q21 0 21 22q0 29-21 47zm-171 81h-42q-22 0-22 21t22 21h42q22 0 22-21t-22-21z"/>`),
		g.Group(children),
	)
}