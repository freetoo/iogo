/*
 *
 * Copyright 2019 The iogo authors.
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
 *
 */

/*
 * contacts: nike: freetoo name: yigui-lu(卢益贵)
 *    wx/qq: 48092788
 *   e-mail: gcode@qq.com
 *      url: https://github.com/freetoo/iogo
 */

package md // import "iogo/unit/md"

import (
	"google.golang.org/grpc/metadata"
)

type MD = metadata.MD

var (
	DecodeKeyValue          = metadata.DecodeKeyValue
	New                     = metadata.New
	Pairs                   = metadata.Pairs
	NewIncomingContext      = metadata.NewIncomingContext
	NewOutgoingContext      = metadata.NewOutgoingContext
	AppendToOutgoingContext = metadata.AppendToOutgoingContext
	FromIncomingContext     = metadata.FromIncomingContext
	FromOutgoingContextRaw  = metadata.FromOutgoingContextRaw
	FromOutgoingContext     = metadata.FromOutgoingContext
)
