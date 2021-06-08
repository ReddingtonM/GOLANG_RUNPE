package help

import "syscall"

var (
	peconv = syscall.MustLoadDLL("run_pe.dll")

	//bool peconv::run_pe(char *payloadPath, char *targetPath)
	procRunPE = peconv.MustFindProc("_ZN6peconv6run_peEPcS0_")
	//BYTE *peconv::load_pe_module(const char *filename, OUT size_t &v_size, bool executable, bool relocate)
	procLoadPEModule1 = peconv.MustFindProc("_ZN6peconv14load_pe_moduleEPKcRybb")
	//BYTE *peconv::load_pe_module(BYTE *dllRawData, size_t r_size, OUT size_t &v_size, bool executable, bool relocate)
	procLoadPEModule2 = peconv.MustFindProc("_ZN6peconv14load_pe_moduleEPhyRybb")
	//WORD peconv::get_nt_hdr_architecture(IN const BYTE *pe_buffer)
	procGetNTHdrArch = peconv.MustFindProc("_ZN6peconv23get_nt_hdr_architectureEPKh")
	//bool peconv::is64bit(IN const BYTE *pe_buffer)
	procIs64Bit = peconv.MustFindProc("_ZN6peconv7is64bitEPKh")
	//bool create_suspended_process(IN LPSTR path, OUT PROCESS_INFORMATION &pi)
	procCreateSuspendedProcess = peconv.MustFindProc("_Z24create_suspended_processPcR20_PROCESS_INFORMATION")
	//bool _run_pe(BYTE *loaded_pe, size_t payloadImageSize, PROCESS_INFORMATION &pi, bool is32bit)
	procRunPE2 = peconv.MustFindProc("_Z7_run_pePhyR20_PROCESS_INFORMATIONb")
	//peconv::ALIGNED_BUF peconv::load_file(IN const char *filename, OUT size_t &read_size)
	procLoadFile = peconv.MustFindProc("_ZN6peconv9load_fileEPKcRy")
	//bool peconv::free_pe_buffer(peconv::ALIGNED_BUF buffer, size_t buffer_size)
	procFreePEBuffer = peconv.MustFindProc("_ZN6peconv14free_pe_bufferEPhy")
	//bool peconv::free_aligned(peconv::ALIGNED_BUF buffer, size_t buffer_size)

	//bool is_target_compatibile(BYTE *payload_buf, size_t payload_size, char *targetPath)
	procIsTargetCompatible = peconv.MustFindProc("_Z21is_target_compatibilePhyPc")
	//WORD peconv::get_subsystem(IN const BYTE* payload)
	procGetSubSystem = peconv.MustFindProc("_ZN6peconv13get_subsystemEPKh")
	//BYTE* peconv::get_nt_hdrs(IN const BYTE *pe_buffer, IN OPTIONAL size_t buffer_size)
	procGetNTHdrs = peconv.MustFindProc("_ZN6peconv11get_nt_hdrsEPKhy")
	//bool peconv::has_relocations(IN const BYTE *pe_buffer)
	procHasRelocations = peconv.MustFindProc("_ZN6peconv15has_relocationsEPKh")
	//ULONGLONG peconv::get_image_base(IN const BYTE *pe_buffer)
	procGetImageBase = peconv.MustFindProc("_ZN6peconv14get_image_baseEPKh")
	//BYTE* peconv::pe_raw_to_virtual(IN const BYTE* payload, IN size_t in_size, OUT size_t &out_size, IN OPTIONAL bool executable,	IN OPTIONAL ULONGLONG desired_base)
	procPERawToVirtual = peconv.MustFindProc("_ZN6peconv17pe_raw_to_virtualEPKhyRyby")
	//bool peconv::relocate_module(IN BYTE* modulePtr, IN SIZE_T moduleSize, IN ULONGLONG newBase, IN ULONGLONG oldBase)
	procRelocateModule = peconv.MustFindProc("_ZN6peconv15relocate_moduleEPhyyy")
	//bool peconv::validate_ptr(IN const void *buffer_bgn, IN SIZE_T buffer_size, IN const void *field_bgn, IN SIZE_T field_size)
	procValidatePtr = peconv.MustFindProc("_ZN6peconv12validate_ptrEPKvyS1_y")
	//bool apply_relocations(PVOID modulePtr, SIZE_T moduleSize, ULONGLONG newBase, ULONGLONG oldBase)
	procApplyRelocations = peconv.MustFindProc("_Z17apply_relocationsPvyyy")
	//bool sections_raw_to_virtual(IN const BYTE *payload, IN SIZE_T payloadSize, OUT BYTE *destBuffer, IN SIZE_T destBufferSize)
	procSectionsRawToVirtual = peconv.MustFindProc("_Z23sections_raw_to_virtualPKhyPhy")
	//bool redirect_to_payload(BYTE *loaded_pe, PVOID load_base, PROCESS_INFORMATION &pi, bool is32bit)
	procRedirectToPayload = peconv.MustFindProc("_Z19redirect_to_payloadPhPvR20_PROCESS_INFORMATIONb")
	//ULONGLONG get_remote_peb_addr(PROCESS_INFORMATION &pi, bool is32bit)
	procGetRemotePebAddr = peconv.MustFindProc("_Z19get_remote_peb_addrR20_PROCESS_INFORMATIONb")
	//BOOL update_remote_entry_point(PROCESS_INFORMATION &pi, ULONGLONG entry_point_va, bool is32bit)
	procUpdateRemoteEntryPoint = peconv.MustFindProc("_Z25update_remote_entry_pointR20_PROCESS_INFORMATIONyb")
	//DWORD peconv::get_entry_point_rva(IN const BYTE *pe_buffer)
	procGetEntryPointRVA = peconv.MustFindProc("_ZN6peconv19get_entry_point_rvaEPKh")
	//bool    process_reloc_block(BASE_RELOCATION_ENTRY *block, SIZE_T entriesNum, DWORD page, PVOID modulePtr, SIZE_T moduleSize, bool is64bit, RelocBlockCallback *callback)
	procProcessRelocBlock = peconv.MustFindProc("_Z19process_reloc_blockPN6peconv22_BASE_RELOCATION_ENTRYEymPvybyy")
)
