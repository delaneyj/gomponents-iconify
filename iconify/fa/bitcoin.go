package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitcoin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1135 512q18 182-131 258q117 28 175 103t45 214q-7 71-32.5 125t-64.5 89t-97 58.5t-121.5 34.5t-145.5 15v255H609v-251q-80 0-122-1v252H333v-255q-18 0-54-.5t-55-.5H24l31-183h111q50 0 58-51V772h16q-6-1-16-1V484q-13-68-89-68H24V252l212 1q64 0 97-1V0h154v247q82-2 122-2V0h154v252q79 7 140 22.5t113 45t82.5 78T1135 512zm-215 545q0-36-15-64t-37-46t-57.5-30.5T745 898t-74-9t-69-3t-64.5 1t-47.5 1v338q8 0 37 .5t48 .5t53-1.5t58.5-4t57-8.5t55.5-14t47.5-21t39.5-30t24.5-40t9.5-51zm-71-476q0-33-12.5-58.5t-30.5-42t-48-28t-55-16.5t-61.5-8t-58-2.5t-54 1t-39.5.5v307q5 0 34.5.5t46.5 0t50-2t55-5.5t51.5-11t48.5-18.5t37-27t27-38.5t9-51z"/>`),
		g.Group(children),
	)
}