package cosmos

type UDF struct {
	client Client
	coll   Collection
	udfID  string
}

type UDFs struct {
	client Client
	coll   Collection
}

func newUDF(coll Collection, udfID string) *UDF {
	coll.client.path = coll.client.path + "/udfs/" + udfID
	coll.client.rType = "udfs"
	coll.client.rLink = coll.client.path
	udf := &UDF{
		client: coll.client,
		coll:   coll,
		udfID:  udfID,
	}

	return udf
}

func newUDFs(coll Collection) *UDFs {
	coll.client.path += "/udfs"
	coll.client.rType = "udfs"
	udfs := &UDFs{
		client: coll.client,
		coll:   coll,
	}

	return udfs
}

func (u *UDFs) Create(newUDF *UDFDefinition, opts ...CallOption) (*UDFDefinition, error) {
	createdUDFResp := &UDFDefinition{}
	_, err := u.client.create(newUDF, &createdUDFResp, opts...)
	if err != nil {
		return nil, err
	}

	return createdUDFResp, err
}

func (u *UDF) Replace(newUDF *UDFDefinition, opts ...CallOption) (*UDFDefinition, error) {
	updatedUDF := &UDFDefinition{}
	_, err := u.client.replace(newUDF, &updatedUDF, opts...)
	if err != nil {
		return nil, err
	}

	return updatedUDF, err
}

func (u *UDFs) ReadAll(opts ...CallOption) ([]UDFDefinition, error) {
	data := struct {
		Udfs  []UDFDefinition `json:"UserDefinedFunctions,omitempty"`
		Count int             `json:"_count,omitempty"`
	}{}
	_, err := u.client.read(&data, opts...)
	if err != nil {
		return nil, err
	}

	return data.Udfs, err
}

func (u *UDF) Delete(opts ...CallOption) (*Response, error) {
	return u.client.delete(opts...)
}
