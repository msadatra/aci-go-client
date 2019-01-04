package models


import (
	"fmt"
	"strconv"

	"github.com/ciscoecosystem/aci-go-client/container"
)

const VmmdompClassName = "vmmDomP"

type VMMDomain struct {
	BaseAttributes
    VMMDomainAttributes 
}
  
type VMMDomainAttributes struct {
    AccessMode       string `json:",omitempty"`
    ArpLearning       string `json:",omitempty"`
    CtrlKnob       string `json:",omitempty"`
    Delimiter       string `json:",omitempty"`
    EnableAVE       string `json:",omitempty"`
    EncapMode       string `json:",omitempty"`
    EnfPref       string `json:",omitempty"`
    EpRetTime       string `json:",omitempty"`
    McastAddr       string `json:",omitempty"`
    Mode       string `json:",omitempty"`
    NameAlias       string `json:",omitempty"`
    PrefEncapMode       string `json:",omitempty"`
    
}
   

func NewVMMDomain(vmmDomPRn, parentDn, description string, vmmDomPattr VMMDomainAttributes) *VMMDomain {
	dn := fmt.Sprintf("%s/%s", parentDn, vmmDomPRn)  
	return &VMMDomain{
		BaseAttributes: BaseAttributes{
			DistinguishedName: dn,
			Description:       description,
			Status:            "created, modified",
			ClassName:         VmmdompClassName,
			Rn:                vmmDomPRn,
		},
        
		VMMDomainAttributes: vmmDomPattr,
         
	}
}

func (vmmDomP *VMMDomain) ToMap() (map[string]string, error) {
	vmmDomPMap, err := vmmDomP.BaseAttributes.ToMap()
	if err != nil {
		return nil, err
	}

    A(vmmDomPMap, "accessMode",vmmDomP.AccessMode)
    A(vmmDomPMap, "arpLearning",vmmDomP.ArpLearning)
    A(vmmDomPMap, "ctrlKnob",vmmDomP.CtrlKnob)
    A(vmmDomPMap, "delimiter",vmmDomP.Delimiter)
    A(vmmDomPMap, "enableAVE",vmmDomP.EnableAVE)
    A(vmmDomPMap, "encapMode",vmmDomP.EncapMode)
    A(vmmDomPMap, "enfPref",vmmDomP.EnfPref)
    A(vmmDomPMap, "epRetTime",vmmDomP.EpRetTime)
    A(vmmDomPMap, "mcastAddr",vmmDomP.McastAddr)
    A(vmmDomPMap, "mode",vmmDomP.Mode)
    A(vmmDomPMap, "nameAlias",vmmDomP.NameAlias)
    A(vmmDomPMap, "prefEncapMode",vmmDomP.PrefEncapMode)
    
	

	return vmmDomPMap, err
}

func VMMDomainFromContainerList(cont *container.Container, index int) *VMMDomain {

	VMMDomainCont := cont.S("imdata").Index(index).S(VmmdompClassName, "attributes")
	return &VMMDomain{
		BaseAttributes{
			DistinguishedName: G(VMMDomainCont, "dn"),
			Description:       G(VMMDomainCont, "descr"),
			Status:            G(VMMDomainCont, "status"),
			ClassName:         VmmdompClassName,
			Rn:                G(VMMDomainCont, "rn"),
		},
        
		VMMDomainAttributes{
        AccessMode : G(VMMDomainCont, "accessMode"),
        ArpLearning : G(VMMDomainCont, "arpLearning"),
        CtrlKnob : G(VMMDomainCont, "ctrlKnob"),
        Delimiter : G(VMMDomainCont, "delimiter"),
        EnableAVE : G(VMMDomainCont, "enableAVE"),
        EncapMode : G(VMMDomainCont, "encapMode"),
        EnfPref : G(VMMDomainCont, "enfPref"),
        EpRetTime : G(VMMDomainCont, "epRetTime"),
        McastAddr : G(VMMDomainCont, "mcastAddr"),
        Mode : G(VMMDomainCont, "mode"),
        NameAlias : G(VMMDomainCont, "nameAlias"),
        PrefEncapMode : G(VMMDomainCont, "prefEncapMode"),
        		
        },
        
	}
}

func VMMDomainFromContainer(cont *container.Container) *VMMDomain {

	return VMMDomainFromContainerList(cont, 0)
}

func VMMDomainListFromContainer(cont *container.Container) []*VMMDomain {
	length, _ := strconv.Atoi(G(cont, "totalCount"))

	arr := make([]*VMMDomain, length)

	for i := 0; i < length; i++ {

		arr[i] = VMMDomainFromContainerList(cont, i)
	}

	return arr
}