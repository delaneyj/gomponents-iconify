package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Walk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M640 0q-66 0-113 47t-47 113t47 113t113 47t113-47t47-113t-47-113T640 0zM275 1616l169-277q28-46 42-92l58-192l175 183l153 367q11 27 34.5 43t53.5 16q40 0 68-28t28-68q0-17-7-36l-140-345q-13-33-20.5-44.5T855 1103L658 891l60-282l56 109q16 34 44 49q6 3 207 105q48 24 63 24q28 0 46-18t18-46q0-41-38-61L885 649l-90-206q-25-55-66-60l-229-26q-32-3-86 19l-242 96q-20 8-39 27t-25 39Q29 742 4 810q-4 9-4 22q0 26 19 45t45 19q20 0 40.5-9.5T131 862l102-268l210-79l-147 687l-188 320q-12 23-12 46q0 40 28 68t68 28q54 0 83-48z"/>`),
		g.Group(children),
	)
}