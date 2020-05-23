package chapter4

import "testing"

func Test_hasRouteBetweenNodes(t *testing.T) {
	t.Run("japan", func(t *testing.T) {
		tokyo := &node{name: "tokyo"}
		kawasaki := &node{name: "kawasaki"}
		saitama := &node{name: "saitama"}
		chiba := &node{name: "chiba"}
		yokohama := &node{name: "yokohama"}
		utsunomiya := &node{name: "utsunomiya"}
		nagoya := &node{name: "nagoya"}
		osaka := &node{name: "osaka"}
		fukushima := &node{name: "fukushima"}
		kyoto := &node{name: "kyoto"}
		shizuoka := &node{name: "shizuoka"}
		kobe := &node{name: "kobe"}
		sendai := &node{name: "sendai"}
		niigata := &node{name: "niigata"}
		obama := &node{name: "obama"}
		takasaki := &node{name: "takasaki"}
		wakayama := &node{name: "wakayama"}

		tokyo.adjecent = []*node{yokohama, saitama, chiba, kawasaki}
		kawasaki.adjecent = []*node{chiba}
		saitama.adjecent = []*node{takasaki, utsunomiya}
		takasaki.adjecent = []*node{utsunomiya}
		utsunomiya.adjecent = []*node{fukushima, sendai, niigata}
		yokohama.adjecent = []*node{kawasaki, shizuoka}
		shizuoka.adjecent = []*node{nagoya}
		nagoya.adjecent = []*node{obama, kyoto}
		kyoto.adjecent = []*node{obama, osaka}
		osaka.adjecent = []*node{kobe, wakayama}

		if !hasRouteBetweenNodes(tokyo, osaka) {
			t.Errorf("tokyo and osaka must have a route!")
		}
	})

	t.Run("mongolia", func(t *testing.T) {
		ulaanbaatar := &node{name: "ulaanbaatar"}
		darkhan := &node{name: "darkhan"}
		erdenet := &node{name: "erdenet"}
		zuunmod := &node{name: "zuunmod"}
		orkhon := &node{name: "orkhon"}
		sukhbaatar := &node{name: "sukhbaatar"}
		shaamar := &node{name: "shaamar"}

		ulaanbaatar.adjecent = []*node{darkhan, zuunmod}
		darkhan.adjecent = []*node{erdenet, orkhon}
		sukhbaatar.adjecent = []*node{orkhon, shaamar}

		if hasRouteBetweenNodes(ulaanbaatar, sukhbaatar) {
			t.Errorf("ulaanbaatar and sukhbaatar do not have a route!")
		}
	})
	t.Run("singapore", func(t *testing.T) {
		singapore := &node{name: "singapore"}
		kualaLumpur := &node{name: "kuala lumpur"}

		singapore.adjecent = []*node{kualaLumpur}

		if !hasRouteBetweenNodes(singapore, singapore) {
			t.Errorf("singapore is singapore!")
		}
	})
}
