package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func View(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 768q79 0 149 30t122 82t83 123t30 149q0 80-30 149t-82 122t-123 83t-149 30q-80 0-149-30t-122-82t-83-122t-30-150q0-79 30-149t82-122t122-83t150-30zm0 640q53 0 99-20t82-55t55-81t20-100q0-53-20-99t-55-82t-81-55t-100-20q-53 0-99 20t-82 55t-55 81t-20 100q0 53 20 99t55 82t81 55t100 20zm0-1152q143 0 284 35t266 105t226 170t166 234q40 83 61 171t21 181h-128q0-118-36-221t-99-188t-150-152t-185-113t-209-70t-217-24q-108 0-217 24t-208 70t-186 113t-149 152t-100 188t-36 221H0q0-92 21-180t61-172q64-132 165-233t227-171t266-104t284-36z"/>`),
		g.Group(children),
	)
}