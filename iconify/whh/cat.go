package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 512q0 32 32 128t32 128q0 65-19.5 123t-57 95.5T736 1024H288q-13 0-22.5-9.5T256 992t9.5-22.5T288 960q14 0 23-9t9-23q0-26-20-60.5t-48-71t-56-75.5t-48-86t-20-91q0-75 60-124q-12-36-60-36q-29 0-60.5-20.5t-49.5-43T0 288q0-13 9.5-22.5T32 256q8 0 17-14t20.5-35T96 164.5t40-40T192 99V0q23 0 54 35.5t40 73.5q48 20 73 69.5T384 288q0 29 12.5 55t35.5 46.5t37.5 31T506 444q4 3 6 4q58 33 96 64q69 57 114.5 143T768 800q0 4 1 9t9 14t22 9t23-15.5t9-48.5q0-35-13-70t-32-66.5t-38-65.5t-32-81.5T704 384q0-48 40-88t88-40q27 0 45.5 18.5T896 320q0 16-6 28.5t-15.5 23T864 384q-32 53-32 128z"/>`),
		g.Group(children),
	)
}