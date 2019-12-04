package bus

// var ParamTypeMap = map[MessageKey]interface{}{}

// func (m *Message) MarshalBSON() ([]byte, error) {
// 	result, err := bson.Marshal(m.Params)
// 	if err != nil {
// 		return []byte{}, err
// 	}

// 	return bson.MarshalAppend(result, m.Key)
// }

// func (m *Message) UnmarshalBSON(raw []byte) error {
// 	params := ParamTypeMap[m.Key]
// 	if params == nil {
// 		return fmt.Errorf("Parameter type not registered")
// 	}

// 	err := bson.Unmarshal(raw, &params)
// 	fmt.Println(params)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func RegisterMessageParamType(key MessageKey, params interface{}) {
// 	ParamTypeMap[key] = params
// }
