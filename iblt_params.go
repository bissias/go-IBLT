package iblt

type IbltParam struct {
    numHashFuncs    int
    itemOverhead    float64
}

var ibltParamMap = map[int]IbltParam{
    1: IbltParam{numHashFuncs: 4, itemOverhead: 1.5},
}
