package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoNotWash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M66 359q3 16 15 30l25-38v-2L72 185l6 6q28 28 58 28v-43q-13 0-28-15q-20-22-47-28L51 86q-1-9-9-14t-17-3q-9 1-14 9T8 95zm209 9h-22l-30 43h82zm117-192q13 0 28 15q17 17 32 23l-30 137v2l25 38q11-10 15-29l58-267q2-9-3-17t-14-9q-9-2-17 3t-9 14l-17 84q-4-6-10-9q-1-1-8.5-9t-14-12.5T413 133zm-218-26l26 35l38 55l-89 128l-30 43l-21 32q-12 17 4 30q2 0 6 1t7 1q6 0 17-9l38-55l30-43l64-92l64 92l30 43l38 55q6 9 17 9q6 0 13-5q8-5 9.5-13t-5.5-16l-21-32l-30-43l-89-128l15-21l53-77l72-103q12-17-4-30q-2 0-6-1t-7-1q-11 0-17 9L287 167l-23 37l-49-71l-83-119q-6-9-17-9q-6 0-13 5t-8 13t4 16z"/>`),
		g.Group(children),
	)
}