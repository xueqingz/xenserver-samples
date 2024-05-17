/*
 * Copyright (c) Cloud Software Group, Inc.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *
 *   1) Redistributions of source code must retain the above copyright
 *      notice, this list of conditions and the following disclaimer.
 *
 *   2) Redistributions in binary form must reproduce the above
 *      copyright notice, this list of conditions and the following
 *      disclaimer in the documentation and/or other materials
 *      provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS
 * FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE
 * COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
 * INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
 * STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED
 * OF THE POSSIBILITY OF SUCH DAMAGE.
 */

package com.xensource.xenapi.samples;

import java.util.Date;
import java.util.Map;

import com.xensource.xenapi.Network;
import com.xensource.xenapi.SR;
import com.xensource.xenapi.Types;
import com.xensource.xenapi.VBD;
import com.xensource.xenapi.VIF;
import com.xensource.xenapi.VM;

/**
 * Creates a VM on the default SR with a network and DVD drive.
 */
public class CreateVM extends TestBase {
    public String getTestName() {
        return "CreateVM";
    }

    protected void TestCore() throws Exception {
        /*First check we can start an HVM on the coordinator*/
        checkCoordinatorHvmCapable();

        VM template = getFirstWindowsTemplate();
        log("Template found: " + template.getNameLabel(connection));

        /* Clone the template */
        String vmName = new Date() + " (made by CreateVM.java)";
        VM newVm = template.createClone(connection, vmName);
        log("New clone: " + newVm.getNameLabel(connection));

        /* Find a storage repository */
        SR theSR = getStorage();
        log("Found SR: " + theSR.getNameLabel(connection));

        /* Find a network */
        Network network = getFirstNetwork();
        log("Network chosen: " + network.getNameLabel(connection));

        /*
         * We have our clone and our network, attach them to each other with a
         * VIF
         */
        makeVIF(newVm, network, "0");

        /* Put the SR uuid into the provision XML */
        Map<String, String> otherConfig = newVm.getOtherConfig(connection);
        String disks = otherConfig.get("disks");
        disks = disks.replace("sr=\"\"", "sr=\"" + theSR.getUuid(connection) + "\"");
        otherConfig.put("disks", disks);
        newVm.setOtherConfig(connection, otherConfig);

        makeCDDrive(newVm);

        /* Now provision the disks */
        log("provisioning... ");
        newVm.provision(connection);
        log("provisioned");

        /* Should have done the trick. Let's see if it starts. */
        log("Starting new VM.....");
        newVm.start(connection, false, false);

        log("Shutting it down (hard).....");
        newVm.hardShutdown(connection);

        log("Destroy VM.....");
        newVm.destroy(connection);
    }

    /*
     * Create a VIF by making a VIF.record and then filling in the necessary
     * fields
     */
    private void makeVIF(VM newVm, Network network, String device) throws Exception {
        VIF.Record newVifRecord = new VIF.Record();

        // These three parameters are used in the command line VIF creation
        newVifRecord.VM = newVm;
        newVifRecord.network = network;
        newVifRecord.device = device;
        newVifRecord.MTU = 1500L;
        newVifRecord.lockingMode = Types.VifLockingMode.NETWORK_DEFAULT;

        VIF.create(connection, newVifRecord);
    }

    private void makeCDDrive(VM vm) throws Exception {
        VBD.Record vbdRecord = new VBD.Record();

        vbdRecord.VM = vm;
        vbdRecord.VDI = null;
        vbdRecord.userdevice = "3";
        vbdRecord.mode = Types.VbdMode.RO;
        vbdRecord.type = Types.VbdType.CD;
        vbdRecord.empty = true;

        VBD.create(connection, vbdRecord);
    }
}
