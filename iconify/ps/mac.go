package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M43 384h170q-8 32-26.5 58.5T149 469q-9 0-15 6t-6 16q0 9 6 15t15 6h214q9 0 15-6t6-15q0-10-6-16t-15-6q-19 0-37.5-26.5T299 384h170q18 0 30.5-12.5T512 341V43q0-18-12.5-30.5T469 0H43Q25 0 12.5 12.5T0 43v298q0 18 12.5 30.5T43 384zm170 85q36-49 43-85q15 60 43 85h-86zM43 43h426v213H43V43zm0 256h426v42H43v-42zm234-86q0 9-6 15.5t-15 6.5t-15-6.5t-6-15.5q0-8 6-14.5t15-6.5t15 6.5t6 14.5zm-85 0q0 9-6.5 15.5T171 235q-9 0-15.5-6.5T149 213q0-8 6.5-14.5T171 192q8 0 14.5 6.5T192 213zm171 0q0 9-6.5 15.5T341 235q-8 0-14.5-6.5T320 213q0-8 6.5-14.5T341 192q9 0 15.5 6.5T363 213z"/>`),
		g.Group(children),
	)
}