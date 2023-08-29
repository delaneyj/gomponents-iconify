package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bonnet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M432 341v-42q0-82-55-137q-24-23-60-40q8-22 8-37q0-35-25-60T240 0t-60 25t-25 60q0 15 8 37q-52 24-83.5 72T48 299v42q-17 0-30 13T5 384v85q0 18 13 30.5T48 512h384q17 0 30-12.5t13-30.5v-85q0-17-13-30t-30-13zM240 43q17 0 30 12.5T283 85q0 14-9 26q-22-4-34-4t-34 4q-9-12-9-26q0-17 13-29.5T240 43zM91 299q0-52 31.5-92t81.5-53q35-11 74 0q34 9 69 38q42 44 42 107v42H91v-42zm298 170h-42v-85h42v85zm-256 0H91v-85h42v85zm22-85h42v85h-42v-85zm64 0h42v85h-42v-85zm64 0h42v85h-42v-85zm-235 0h21v85H48v-85zm363 85v-85h21v85h-21z"/>`),
		g.Group(children),
	)
}