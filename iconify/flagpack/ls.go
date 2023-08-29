package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackLs0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackLs1)"><use href="#flagpackLs0"/><path fill="#55BA07" fill-rule="evenodd" d="M0 16h32v8H0v-8Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="M0 8h32v8H0V8Z" clip-rule="evenodd"/><path fill="#3D58DB" fill-rule="evenodd" d="M0 0h32v8H0V0Z" clip-rule="evenodd"/><path fill="#1D1D1D" fill-rule="evenodd" d="M15.25 8.24c-.523.14-.799.502-.799 1.141c0 .73.359 1.416.799 1.77V8.24Zm1.25 2.89c.428-.357.772-1.03.772-1.747c0-.62-.272-.98-.772-1.13v2.878Zm-.042-3.398c.818.182 1.41.74 1.41 1.645c0 .659-.313 1.46-.798 2.009l2.433 2.387l.6-.198l.897 1.36S18.734 16 15.933 16C13.133 16 11 14.935 11 14.935l.763-1.162l.617.202l2.257-2.548c-.507-.55-.836-1.375-.836-2.05c0-.938.635-1.502 1.498-1.663a.625.625 0 0 1 1.159.018Zm-4.456 6.564l-.254.353h.494l-.24-.353Zm.289.492l-.436.088l.436.152v-.24Zm.222-.508l.223.52l.553-.15l-.11-.21l-.666-.16Zm1.204 1.038l-.813-.16l-.093-.157l.613-.133l.293.45Zm.047-.745l.253.455l.555-.17l-.154-.203l-.654-.082Zm.456.836l-.1-.188l.615-.165l.354.433l-.87-.08Zm1.148-.313l.792-.209l-.098-.18h-1.028l.334.39Zm.284.401l-.126-.195l.76-.249l.404.444h-1.038Zm.963-.776l.243.298l.659-.228l-.083-.156l-.82.086Zm1.43.62l-.886.098l-.13-.22l.67-.22l.346.342Zm-.057-.764l.249.305l.594-.305l-.165-.092l-.678.092Zm.515.682l-.108-.165l.566-.34l.288.34l-.746.165Zm.562-.95l.36.344l.349-.343l-.11-.14l-.6.14Zm.57.461l.292-.29l.211.29l-.504.162v-.162Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackLs1"><use href="#flagpackLs0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}