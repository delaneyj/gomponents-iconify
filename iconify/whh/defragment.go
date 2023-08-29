package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Defragment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 192V64h128v128H896zm64 192H832V256h128v128zM704 0h128v128H704V0zm0 256v128h-64v128H512V384h64V256h128zm-64 448h128v-64H640V512h128v-64h128v448q0 53-37.5 90.5T768 1024H256q-53 0-90.5-37.5T128 896v-64h128v-64H128V640h128v64h256v-64h128v64zM512 832h128v-64H512v64zM384 640v-64H256V448h128v64h128v128H384zm64-256H320V256h128v128zm0-320h128v128H448V64zM192 0h128v128H192V0zm0 320H64V192h128v128zm-64 256H0V448h128v128z"/>`),
		g.Group(children),
	)
}