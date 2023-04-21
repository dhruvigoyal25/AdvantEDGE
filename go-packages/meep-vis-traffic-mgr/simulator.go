/*
 * Copyright (c) 2022  The AdvantEDGE Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package vistrafficmgr

import (
	"encoding/hex"
	//"errors"
	"math/rand"
	"time"

	log "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-logger"
)

type message_broker_simu struct {
	running              bool
	simulatedV2xMessages map[int32][]byte
	simulateBrokerTicker *time.Ticker
}

func (simu *message_broker_simu) Init(tm *TrafficMgr) (err error) {
	log.Info("message_broker_simu: Init")

	simu.running = false
	const (
		DENM              int32 = 1
		CAM               int32 = 2
		POI               int32 = 3
		SPATEM            int32 = 4
		MAPEM             int32 = 5
		IVIM              int32 = 6
		EV_RSR            int32 = 7
		TISTPGTRANSACTION int32 = 8
		SREM              int32 = 9
		SSEM              int32 = 10
		EVCSN             int32 = 11
		SAEM              int32 = 12
		RTCMEM            int32 = 13
	)
	simu.simulatedV2xMessages = map[int32][]byte{
		CAM: {
			0x03, 0x12, 0x00, 0x05, 0x01, 0x03, 0x81, 0x00, /* ........ */
			0x40, 0x03, 0x80, 0x65, 0x20, 0x50, 0x02, 0x80, /* @..e P.. */
			0x00, 0x41, 0x01, 0x00, 0x14, 0x00, 0x04, 0xe5, /* .A...... */
			0x48, 0x14, 0x00, 0x72, 0x0c, 0x65, 0x09, 0xbf, /* H..r.e.. */
			0x1e, 0xb7, 0x81, 0xdf, 0x08, 0x4d, 0x55, 0x4e, /* .....MUN */
			0x80, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
			0x07, 0xd1, 0x00, 0x00, 0x02, 0x02, 0x48, 0x14, /* ......H. */
			0x00, 0x72, 0x09, 0xbf, 0x40, 0x5a, 0x8b, 0x8d, /* .r..@Z.. */
			0x5b, 0xee, 0x72, 0xe4, 0xe9, 0xc1, 0x90, 0x19, /* [.r..... */
			0x03, 0x84, 0x39, 0x29, 0xd0, 0x40, 0xe1, 0x1f, /* ..9).@.. */
			0xc0, 0x00, 0x7e, 0x82, 0x98, 0x8a, 0x83, 0x37, /* ..~....7 */
			0xfe, 0xeb, 0xff, 0xf6, 0x00, 0x00, 0x00, 0x2b, /* .......+ */
			0xff, 0xd0, 0xff, 0xef, 0x32, 0xbe, 0x7f, 0xff, /* ....2... */
			0x60, 0x05, 0x77, 0xff, 0xfd, 0x8b, 0x53, 0xff, /* `.w...S. */
			0xf8, 0xc0, 0x01, 0x24, 0x00, 0x02, 0x0b, 0xa0, /* ...$.... */
			0x6a, 0xaf, 0x23, 0x01, 0x02, 0x05, 0x80, 0x05, /* j.#..... */
			0x01, 0x01, 0x81, 0xa8, 0x43, 0x81, 0x01, 0x01, /* ....C... */
			0x80, 0x03, 0x00, 0x80, 0x31, 0x63, 0xf3, 0x55, /* ....1c.U */
			0x36, 0xc8, 0x3e, 0xd0, 0x10, 0x83, 0x00, 0x00, /* 6.>..... */
			0x00, 0x00, 0x00, 0x22, 0x50, 0x26, 0xe0, 0x84, /* ..."P&.. */
			0x00, 0x24, 0x01, 0x02, 0x80, 0x01, 0x24, 0x81, /* .$....$. */
			0x04, 0x03, 0x01, 0xff, 0xfc, 0x80, 0x01, 0x25, /* .......% */
			0x81, 0x05, 0x04, 0x01, 0xff, 0xff, 0xff, 0x80, /* ........ */
			0x80, 0x83, 0x46, 0xf7, 0x82, 0xd1, 0x33, 0xdb, /* ..F...3. */
			0x96, 0xbb, 0xac, 0x7d, 0xdf, 0x85, 0xc0, 0xb2, /* ...}.... */
			0x33, 0x6c, 0x4a, 0xb8, 0x39, 0x01, 0x1d, 0x67, /* 3lJ.9..g */
			0xb1, 0x00, 0x48, 0x77, 0x4a, 0x15, 0xbf, 0x7f, /* ..HwJ... */
			0x45, 0xeb, 0x80, 0x80, 0xd3, 0x74, 0x1a, 0xc0, /* E....t.. */
			0x58, 0x6d, 0x85, 0xac, 0x63, 0x99, 0xee, 0x3c, /* Xm..c..< */
			0x60, 0x20, 0xa2, 0x55, 0xe1, 0x88, 0x6d, 0x53, /* ` .U..mS */
			0xf4, 0x32, 0xf3, 0xea, 0xe9, 0x52, 0x13, 0x31, /* .2...R.1 */
			0x4c, 0x3d, 0x52, 0x44, 0xd6, 0x5f, 0xf7, 0x08, /* L=RD._.. */
			0x3c, 0xe7, 0x97, 0xc7, 0x0f, 0xb3, 0x50, 0xc9, /* <.....P. */
			0xee, 0xf1, 0xd3, 0x95, 0x76, 0x80, 0x9f, 0x8b, /* ....v... */
			0x6c, 0xb2, 0x31, 0xe6, 0x98, 0xfe, 0xa9, 0xbf, /* l.1..... */
			0xfc, 0x80, 0xc8, 0x83, 0x80, 0x80, 0x63, 0xca, /* ......c. */
			0xdc, 0x41, 0xab, 0x68, 0x17, 0x41, 0xf0, 0xd0, /* .A.h.A.. */
			0xf1, 0xea, 0xab, 0x6e, 0x6c, 0xda, 0xa8, 0x66, /* ...nl..f */
			0x23, 0x21, 0xb3, 0xb2, 0xc5, 0x76, 0x49, 0xfa, /* #!...vI. */
			0xe6, 0xc7, 0x1f, 0xcf, 0xa4, 0xc6, 0x74, 0x5d, /* ......t] */
			0x2d, 0x32, 0x69, 0x72, 0x29, 0xbb, 0xbf, 0xb6, /* -2ir)... */
			0x22, 0xb6, 0x60, 0xe6, 0xb5, 0xcd, 0x31, 0x9a, /* ".`...1. */
			0x97, 0xde, 0x0d, 0xdf, 0xf4, 0xb1, 0x8a, 0x0f, /* ........ */
			0x65, 0xc6, 0x53, 0xf0, 0xd6, 0x5d,
		},
		DENM: {
			0x03, 0x12, 0x00, 0xf1, 0x01, 0x03, 0x81, 0x00, /* ........ */
			0x40, 0x03, 0x80, 0x81, 0x8c, 0x20, 0x40, 0x01, /* @.... @. */
			0x00, 0x00, 0x58, 0x02, 0x00, 0x1e, 0xe6, 0x00, /* ..X..... */
			0x00, 0x3c, 0x00, 0x04, 0xe5, 0x48, 0x14, 0x00, /* .<...H.. */
			0x72, 0x06, 0x6b, 0x24, 0xd0, 0x1e, 0xb7, 0x81, /* r.k$.... */
			0x49, 0x08, 0x4d, 0x55, 0x71, 0x80, 0x00, 0x00, /* I.MUq... */
			0x00, 0x1e, 0xb7, 0x8b, 0x7b, 0x08, 0x4d, 0xc1, /* ....{.M. */
			0x8b, 0x07, 0xd0, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
			0x00, 0x07, 0xd2, 0x00, 0x00, 0x02, 0x01, 0x71, /* .......q */
			0xd9, 0x5b, 0x9a, 0xc7, 0xb8, 0xec, 0xad, 0xcd, /* .[...... */
			0x00, 0x1b, 0x90, 0xc0, 0xc1, 0xb4, 0xe7, 0xa4, /* ........ */
			0x30, 0x30, 0x36, 0x58, 0x35, 0x45, 0xc7, 0x47, /* 006X5E.G */
			0xb7, 0x39, 0x79, 0x38, 0xbf, 0xff, 0xff, 0xfe, /* .9y8.... */
			0x11, 0xdb, 0xba, 0x1f, 0x0b, 0x01, 0xe8, 0x06, /* ........ */
			0x30, 0xf1, 0xc1, 0x80, 0x25, 0xa7, 0x62, 0x40, /* 0...%.b@ */
			0xa7, 0xf2, 0xea, 0x02, 0x0b, 0x63, 0x38, 0x7f, /* .....c8. */
			0x2e, 0xa0, 0x20, 0xb6, 0x33, 0x87, 0xf2, 0xea, /* .. .3... */
			0x02, 0x0b, 0x63, 0x38, 0x7f, 0x2e, 0xa0, 0x20, /* ..c8...  */
			0xb6, 0x33, 0x87, 0xf2, 0xea, 0x02, 0x0b, 0x63, /* .3.....c */
			0x38, 0x50, 0x01, 0x25, 0x00, 0x02, 0x0b, 0x89, /* 8P.%.... */
			0x12, 0x89, 0x6a, 0x0d, 0x1e, 0xb7, 0x81, 0x49, /* ..j....I */
			0x08, 0x4d, 0x55, 0x71, 0x06, 0xa2, 0x81, 0x01, /* .MUq.... */
			0x01, 0x80, 0x03, 0x00, 0x80, 0x31, 0x63, 0xf3, /* .....1c. */
			0x55, 0x36, 0xc8, 0x3e, 0xd0, 0x10, 0x83, 0x00, /* U6.>.... */
			0x00, 0x00, 0x00, 0x00, 0x22, 0x4e, 0x2c, 0xa0, /* ...."N,. */
			0x84, 0x00, 0x24, 0x01, 0x02, 0x80, 0x01, 0x24, /* ..$....$ */
			0x81, 0x04, 0x03, 0x01, 0xff, 0xfc, 0x80, 0x01, /* ........ */
			0x25, 0x81, 0x05, 0x04, 0x01, 0xff, 0xff, 0xff, /* %....... */
			0x80, 0x80, 0x83, 0x69, 0x9e, 0x54, 0x31, 0x01, /* ...i.T1. */
			0x42, 0x10, 0x24, 0xed, 0xf0, 0x50, 0xb6, 0xf1, /* B.$..P.. */
			0xe3, 0xb7, 0xdd, 0xf7, 0xcc, 0xc6, 0x5c, 0xb3, /* ......\. */
			0x1b, 0x1f, 0x72, 0xec, 0x5d, 0xad, 0x92, 0x4c, /* ..r.]..L */
			0x4a, 0x4e, 0x01, 0x80, 0x80, 0x70, 0xee, 0xfd, /* JN...p.. */
			0x24, 0x44, 0xf4, 0x3f, 0x98, 0x84, 0x78, 0x7d, /* $D.?..x} */
			0x5d, 0xa5, 0x2e, 0x2e, 0x5a, 0x4b, 0x3a, 0xc7, /* ]...ZK:. */
			0x1e, 0x87, 0xc2, 0x54, 0xf6, 0xb8, 0xab, 0xef, /* ...T.... */
			0xef, 0x8a, 0xa2, 0xc0, 0x30, 0xa6, 0x20, 0x1b, /* ....0. . */
			0x45, 0xc1, 0x09, 0x55, 0x94, 0x31, 0x96, 0x07, /* E..U.1.. */
			0xfa, 0x05, 0x3b, 0x92, 0xdd, 0xc7, 0x98, 0xf7, /* ..;..... */
			0x2c, 0xfd, 0xe4, 0x80, 0x0c, 0x97, 0xfa, 0xa3, /* ,....... */
			0xfc, 0x9c, 0x51, 0xaf, 0xdd, 0x80, 0x80, 0x83, /* ..Q..... */
			0xc0, 0x3e, 0x25, 0x01, 0x0f, 0xe9, 0x8f, 0x93, /* .>%..... */
			0x6b, 0x8d, 0x5b, 0xe9, 0x69, 0x5d, 0xf3, 0xa5, /* k.[.i].. */
			0x33, 0x7e, 0x42, 0xa0, 0x8e, 0xcf, 0xf4, 0x3c, /* 3~B....< */
			0x93, 0xe1, 0x6e, 0xe0, 0x01, 0x63, 0x97, 0x30, /* ..n..c.0 */
			0x31, 0x99, 0xe9, 0x61, 0x82, 0x57, 0x86, 0xe7, /* 1..a.W.. */
			0x15, 0xf0, 0x96, 0xcf, 0x7e, 0x27, 0x55, 0x6a, /* ....~'Uj */
			0x3a, 0x0a, 0x0d, 0x9f, 0xd9, 0x35, 0xd8, 0x7a, /* :....5.z */
			0x56, 0x06, 0xcf, 0x81, 0xfe, 0x5e, 0x24,
		},
		IVIM: {0xDE, 0x5E},
		MAPEM: {
			0x03, 0x12, 0x00, 0x50, 0x01, 0x03, 0x81, 0x00, /* ...P.... */
			0x40, 0x03, 0x80, 0x6f, 0x20, 0x40, 0x03, 0x80, /* @..o @.. */
			0x00, 0x3b, 0x01, 0x00, 0xae, 0xd0, 0x00, 0x00, /* .;...... */
			0xbc, 0x00, 0x00, 0x13, 0x95, 0x2a, 0x6b, 0x79, /* .....*ky */
			0x02, 0x86, 0x8e, 0x5e, 0x1e, 0xb7, 0x7e, 0xa0, /* ...^..~. */
			0x08, 0x4d, 0x51, 0xf0, 0x00, 0x00, 0x00, 0x00, /* .MQ..... */
			0x1e, 0xb7, 0x7e, 0xb2, 0x08, 0x4d, 0x51, 0xf2, /* ..~..MQ. */
			0x13, 0x88, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
			0x07, 0xd3, 0x00, 0x00, 0x02, 0x05, 0x00, 0x01, /* ........ */
			0x3d, 0xb0, 0x08, 0x00, 0x00, 0x10, 0x00, 0x10, /* =....... */
			0xbb, 0x90, 0x22, 0xa2, 0xe1, 0xd0, 0x83, 0x9c, /* .."..... */
			0xb4, 0x14, 0x10, 0x09, 0x20, 0x50, 0xa0, 0x00, /* .... P.. */
			0x00, 0x00, 0x00, 0xa1, 0x68, 0x61, 0xa5, 0x1f, /* ....ha.. */
			0x74, 0xc0, 0x81, 0x43, 0x21, 0x01, 0x42, 0x20, /* t..C!.B  */
			0x00, 0x00, 0x00, 0x00, 0xa3, 0x34, 0xc3, 0x4a, /* .....4.J */
			0x86, 0xe4, 0x80, 0x40, 0x01, 0x8a, 0x00, 0x02, /* ...@.... */
			0x0b, 0x79, 0xdd, 0x9c, 0x21, 0x54, 0x80, 0x64, /* .y..!T.d */
			0xc7, 0x8c, 0x07, 0x98, 0x1e, 0x0d, 0x4a, 0x80, /* ......J. */
			0x80, 0xda, 0xfb, 0x20, 0x62, 0xda, 0x36, 0xea, /* ... b.6. */
			0xe3, 0x5a, 0x45, 0x30, 0xfe, 0x5d, 0xc4, 0x60, /* .ZE0.].` */
			0x01, 0x60, 0xa8, 0xb7, 0xff, 0x34, 0x77, 0xbd, /* .`...4w. */
			0x79, 0x9f, 0x17, 0x08, 0xd9, 0x1a, 0x72, 0x28, /* y.....r( */
			0xc7, 0xcd, 0xa2, 0x6c, 0x6e, 0x02, 0xf8, 0x54, /* ...ln..T */
			0xe0, 0xbc, 0x21, 0xa6, 0x86, 0x71, 0x44, 0x1f, /* ..!..qD. */
			0xc4, 0xcd, 0xe6, 0x53, 0xe8, 0x19, 0xb2, 0x90, /* ...S.... */
			0xab, 0x45, 0x04, 0x62, 0xa1, 0xfa, 0xd7, 0x13, /* .E.b.... */
			0x09,
		},
		SPATEM: {
			0x00, 0x04, 0x02, 0x12, 0x00, 0x00, 0x00, 0x00, /* ........ */
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x86, 0xdd, /* ........ */
			0x60, 0x30, 0x00, 0x00, 0x00, 0xc2, 0x11, 0x01, /* `0...... */
			0x80, 0xf8, 0x0f, 0x80, 0xf8, 0x0f, 0x80, 0xf8, /* ........ */
			0x61, 0xa9, 0x74, 0x07, 0xbb, 0xbb, 0x31, 0x82, /* a.t...1. */
			0xff, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, /* ........ */
			0x0a, 0x2a, 0x0a, 0x2a, 0x00, 0xc2, 0x77, 0xe8, /* .*.*..w. */
			0x03, 0x12, 0x00, 0x50, 0x01, 0x03, 0x81, 0x00, /* ...P.... */
			0x40, 0x03, 0x80, 0x58, 0x20, 0x40, 0x03, 0x80, /* @..X @.. */
			0x00, 0x24, 0x01, 0x00, 0xae, 0xd2, 0x00, 0x00, /* .$...... */
			0xbc, 0x00, 0x00, 0x13, 0x95, 0x2a, 0x6b, 0x79, /* .....*ky */
			0x02, 0x86, 0x91, 0x18, 0x1e, 0xb7, 0x7e, 0xa0, /* ......~. */
			0x08, 0x4d, 0x51, 0xf0, 0x00, 0x00, 0x00, 0x00, /* .MQ..... */
			0x1e, 0xb7, 0x7e, 0xb2, 0x08, 0x4d, 0x51, 0xf2, /* ..~..MQ. */
			0x13, 0x88, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, /* ........ */
			0x07, 0xd4, 0x00, 0x00, 0x02, 0x04, 0x00, 0x01, /* ........ */
			0x3d, 0xb0, 0x00, 0x18, 0x80, 0x00, 0x85, 0xdc, /* =....... */
			0x81, 0x00, 0x00, 0x1e, 0x77, 0xfb, 0x3e, 0xe0, /* ....w.>. */
			0x00, 0x32, 0x14, 0x30, 0x46, 0x0f, 0x22, 0xc2, /* .2.0F.". */
			0x30, 0x78, 0x02, 0xa8, 0x40, 0x01, 0x89, 0x00, /* 0x..@... */
			0x02, 0x0b, 0x79, 0xdd, 0xa6, 0xc6, 0x8c, 0x80, /* ..y..... */
			0x64, 0xc7, 0x8c, 0x07, 0x98, 0x1e, 0x0d, 0x4a, /* d......J */
			0x80, 0x80, 0x64, 0x2a, 0x2c, 0x10, 0x68, 0x74, /* ..d*,.ht */
			0x22, 0x00, 0xc3, 0xaa, 0xdb, 0x9b, 0x04, 0x02, /* "....... */
			0x01, 0xb6, 0x41, 0x17, 0x73, 0x5b, 0x54, 0x9a, /* ..A.s[T. */
			0x9c, 0x0f, 0x02, 0xe4, 0x70, 0x7e, 0xd7, 0x9c, /* ....p~.. */
			0xc2, 0xf4, 0xab, 0x82, 0x96, 0x52, 0xb5, 0xa2, /* .....R.. */
			0x91, 0xe4, 0xac, 0x98, 0x83, 0x09, 0x8b, 0xce, /* ........ */
			0xb5, 0xfd, 0xa4, 0xec, 0x12, 0xba, 0xb0, 0x7c, /* .......| */
			0xe4, 0x93, 0xcb, 0x7c, 0x20, 0x91, 0xc5, 0xa3, /* ...| ... */
			0x09, 0x7e,
		},
	}

	return nil
}

func (simu *message_broker_simu) Run(tm *TrafficMgr) (err error) {
	log.Info("message_broker_simu: Run: simu.running: ", simu.running)

	simu.simulateBrokerTicker = time.NewTicker(time.Second)
	simu.running = true
	go func() {
		for range simu.simulateBrokerTicker.C {
			if !simu.running {
				log.Info("message_broker_simu: Run: Exit loop")
				break
			}
			// Generate new V2X message
			i := int32(rand.Intn(2)) + 1
			v2xMessage := simu.simulatedV2xMessages[i]
			tm.v2x_notify(v2xMessage, i, nil, nil)
		} // End of 'for' statement
	}()

	return nil
}

func (simu *message_broker_simu) Stop(tm *TrafficMgr) (err error) {
	log.Info("message_broker_simu: Stop: simu.running: ", simu.running)

	if simu.running {
		simu.running = false
		time.Sleep(2 * time.Second)
		simu.simulateBrokerTicker.Stop()
		simu.simulateBrokerTicker = nil
	}

	return nil
}

func (simu *message_broker_simu) Send(tm *TrafficMgr, msgContent string, msgEncodeFormat string, stdOrganization string, msgType *int32) (err error) {
	log.Info("message_broker_simu: Send")

	if msgEncodeFormat == "hexadump" {
		content, err := hex.DecodeString(msgContent)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		log.Info("message_broker_simu: Send: Publish content : ", content)
		log.Info("message_broker_simu: Send: msgEncodeFormat: ", msgEncodeFormat)
		log.Info("message_broker_simu: Send: stdOrganization: ", stdOrganization)
	}

	return nil
}
